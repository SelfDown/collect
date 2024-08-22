package collect

import (
	"encoding/base64"
	common "github.com/SelfDown/collect/src/collect/common"
	"github.com/SelfDown/collect/src/collect/config"
	cacheHandler "github.com/SelfDown/collect/src/collect/service_imp/cache_handler"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"net"
	"reflect"
	"strings"
)

/*
*
* 请求前处理参数
 */
type BeforeLoader struct {
}

func handlerValueType(template *collect.Template) {
	paramConfig := template.Params
	for name, config := range paramConfig {
		if utils.IsValueEmpty(config.Type) {
			continue
		}
		// 模板渲染值
		value := template.GetParam(name)
		if utils.IsValueEmpty(value) {
			continue
		}
		value = utils.CastValue(value, config.Type)
		// 重新设置值
		template.SetParam(name, value)
	}

}
func handlerValueTemplate(template *collect.Template) {
	paramConfig := template.Params
	//处理template
	for name, config := range paramConfig {
		if utils.IsValueEmpty(config.Template) {
			continue
		}
		paramPool := template.GetParams()
		// 模板渲染值
		value := utils.RenderTplExec(config.TemplateTpl, paramPool, config.Exec)
		// 模板变量赋值
		//template.paramPool[name] = value
		template.SetParam(name, value)
	}
}

func handlerDefaultValue(template *collect.Template) {

	paramConfig := template.Params
	// 处理默认值
	for name, config := range paramConfig {
		//如果没有配置default返回
		if utils.IsValueEmpty(config.Default) {
			continue
		}
		//如果有值则跳过
		if !utils.IsEmpty(name, template.GetParams()) {
			continue
		}
		// 根据名称设置默认值
		//template.paramPool[name] = config.Default
		template.SetParam(name, config.Default)
	}
	//设置当前session用户
	template.SetParam("session_user_id", template.OpUser)
}
func handlerCheckValue(template *collect.Template) *common.Result {
	paramConfig := template.Params
	// 处理默认值
	for name, config := range paramConfig {
		//如果没有配置default返回
		if utils.IsValueEmpty(config.Check.Template) {
			continue
		}
		check := config.Check
		if check.TemplateTpl == nil {
			return common.NotOk(name + "check中【template】模板不存在")
		}
		if check.ErrMsgTpl == nil {
			return common.NotOk(name + "check中【err_msg】模板不存在")
		}
		paramPool := template.GetParams()
		result := utils.RenderTplBool(check.TemplateTpl, paramPool)
		if !result {
			msg := utils.RenderTpl(check.ErrMsgTpl, paramPool)
			return common.NotOk(msg)
		}

	}
	return common.Ok(nil, "成功")
}
func (t *BeforeLoader) HandlerReqParam(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {

	// 处理默认值
	handlerDefaultValue(template)
	//处理数据模板
	handlerValueTemplate(template)
	//处理数据类型
	handlerValueType(template)
	// 检查数据
	checkResult := handlerCheckValue(template)
	if !checkResult.Success {
		return checkResult
	}
	return common.Ok(nil, "参数检查完毕")

}

/**
* 处理参数
 */
func (t *BeforeLoader) HandlerParams(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	return handlerParams(template, template.HandlerParams, ts)
}

/**
* 处理参数
 */
func (t *BeforeLoader) HandlerLoginCheck(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	// 如果是http 请求，并且配置必须登录，如果没有用户ID,则提示必须登录
	mustLoginConfig := utils.GetAppKeyWithDefault("must_login", "true")
	mustLogin := true
	// 总开关
	if mustLoginConfig != "true" {
		mustLogin = false
	}
	// 服务开关
	if template.MustLogin != nil && *template.MustLogin == false {
		mustLogin = false
	}
	opUser := sessioinUserId(*ts.GetSession())
	// 用户登录判断
	if config.IsHttp && mustLogin && utils.IsValueEmpty(opUser) {
		errMsg := "请登录！！！"
		return common.NotOk(errMsg)
	}
	return common.Ok(nil, "检查成功")
}

/*
*
处理Basic Auth 认证
*/
func (t *BeforeLoader) HandlerBasicAuth(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	c := ts.context
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if utils.IsValueEmpty(auth) || utils.IsValueEmpty(auth[0]) {
		return common.Ok(nil, "无需Basic auth 认证处理")
	}
	if len(auth) != 2 || auth[0] != "Basic" {

		return common.NotOk("无Basic 字段")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return common.NotOk("Basic realm=Authorization Required")
	}
	username := pair[0]
	password := pair[1]
	var ret *common.Result
	params := make(map[string]interface{})
	service := utils.GetAppKey("basic_auth_service")
	params["service"] = service
	params["username"] = username
	params["password"] = password
	utils.Block{
		Try: func() {
			ret = ts.Result(params, false)
		},
		Catch: func(e utils.Exception) {
			dv := reflect.ValueOf(e)
			ret = common.NotOk(dv.String())
		},
	}.Do()
	if !ret.Success {
		return ret
	}
	return common.Ok(ret.Data, ret.Msg)
}

func (t *BeforeLoader) HandlerCache(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	handlerParam := template.Cache
	if handlerParam.EnableTpl == nil {
		return common.Ok(nil, "无缓存，无需处理")
	}
	// 获取缓存
	handlerParam.Method = cacheHandler.CacheGetName
	ret := HandlerOneParams(&handlerParam, template, ts)
	if ret.Success && ret.Data != nil {
		ret.SetFinish(true)
	}

	return ret
}

func (t *AfterLoader) HandlerCache(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	handlerParam := template.Cache
	if handlerParam.EnableTpl == nil {
		return common.Ok(nil, "无缓存，无需处理")
	}
	// 设置缓存
	handlerParam.Method = cacheHandler.CacheSetName
	ret := HandlerOneParams(&handlerParam, template, ts)
	return ret
}

// GetLocalIPs 返回本机的所有 IPv4 地址（字符串数组）
func GetLocalIPs() ([]interface{}, error) {
	var ips []interface{}

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// 获取接口的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 跳过非 IPv4 和回环地址
			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			ips = append(ips, ip.String())
		}
	}

	return ips, nil
}

// 处理代理
func (t *BeforeLoader) HandlerProxy(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	handlerParam := template.Proxy
	if handlerParam.EnableTpl == nil {
		return common.Ok(nil, "无代理，无需处理")
	}
	ips, _ := GetLocalIPs()
	params := template.GetParams()
	// 获取IP
	params["local_machine_ips"] = ips
	// 设置代理
	params["proxy_service"] = params["service"]
	ret := HandlerOneParams(&handlerParam, template, ts)
	if ret.Success {
		ret.SetFinish(true)

	}
	return ret
}

// 防止重复请求
func (t *BeforeLoader) PreventDuplication(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	handlerParam := template.PreventDuplication
	if handlerParam.EnableTpl == nil {
		return common.Ok(nil, "重复请求为配置，无需处理")
	}
	ret := HandlerOneParams(&handlerParam, template, ts)
	// 如果设置了缓存
	if ret.Success && gocast.ToString(ret.GetData()) == "1" {
		ret.SetFinish(true)
		ret.Success = false
		ret.Code = "-1"
		ret.Msg = "您重复请求[" + template.GetService() + "]剩余：" + gocast.ToString(ret.Count) + "毫秒"
	}

	return ret
}
