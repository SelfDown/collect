package collect

import (
	common "collect.mod/src/collect/common"
	utils "collect.mod/src/collect/utils"
)

/**
* 请求前处理参数

 */
type BeforeLoader struct {
}

func handlerValueType(template Template) {
	paramConfig := template.Params
	for name, config := range paramConfig {
		if utils.IsValueEmpty(config.Type) {
			continue
		}
		// 模板渲染值
		value := template.param_pool[name]
		if utils.IsValueEmpty(value) {
			continue
		}
		value = utils.CastValue(value, config.Type)
		//switch config.Type {
		//case "int":
		//	value = gocast.ToInt(value)
		//	break
		//case "bool":
		//	value = gocast.ToBool(value)
		//	break
		//case "float":
		//	value = gocast.ToFloat(value)
		//	break
		//}
		// 重新设置值
		template.param_pool[name] = value
	}
}
func handlerValueTemplate(template Template) {
	paramConfig := template.Params
	//处理template
	for name, config := range paramConfig {
		if utils.IsValueEmpty(config.Template) {
			continue
		}
		// 模板渲染值
		value := utils.RenderTplExec(config.TemplateTpl, template.param_pool, config.Exec)
		// 模板变量赋值
		template.param_pool[name] = value
	}
}

func handlerDefaultValue(template Template) {

	paramConfig := template.Params
	// 处理默认值
	for name, config := range paramConfig {
		//如果没有配置default返回
		if utils.IsValueEmpty(config.Default) {
			continue
		}
		//如果有值则跳过
		if !utils.IsEmpty(name, template.param_pool) {
			continue
		}
		// 根据名称设置默认值
		template.param_pool[name] = config.Default
	}
}
func handlerCheckValue(template Template) *common.Result {
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
		result := utils.RenderTplBool(check.TemplateTpl, template.param_pool)
		if !result {
			msg := utils.RenderTpl(check.ErrMsgTpl, template.param_pool)
			return common.NotOk(msg)
		}

	}
	return common.Ok(nil, "成功")
}
func (t *BeforeLoader) HandlerReqParam(config Plugin, template Template, routerAll *RouterAll) *common.Result {

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

func (t *BeforeLoader) HandlerParams(config Plugin, template Template, routerAll *RouterAll) *common.Result {
	return handlerParams(template, template.HandlerParams)
}
