package collect

import (
	common "collect.mod/src/collect/common"
	"collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

/**
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

func (t *BeforeLoader) HandlerParams(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {
	return handlerParams(template, template.HandlerParams, ts)
}
