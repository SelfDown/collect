package collect

import (
	common "collect.mod/src/collect/common"
	"collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type AfterLoader struct {
}

/**
处理handler param 中的节点
*/
func HandlerOneParams(handlerParam *collect.HandlerParam,
	template *collect.Template,
	ts *TemplateService) *common.Result {
	params := template.GetParams()
	// 处理是否启用
	if !utils.IsValueEmpty(handlerParam.Enable) { // 如果配置enable ，则判断是否启用
		enable := utils.RenderTplDataBool(handlerParam.EnableTpl, params)
		if !enable { // 如果结果为false 则跳过处理器
			return common.Ok(nil, "跳过执行")
		}

	}
	// 获取数据处理模块
	module := GetModuleRegister(handlerParam.Key)
	if module == nil {
		return common.NotOk(handlerParam.Key + "处理模块没有找到")
	}
	//数据处理
	r := module.HandlerData(template, handlerParam, ts)
	if !r.Success { // 如果结果不成功直接返回
		return r
	}
	//todo 这里处理count
	saveField := handlerParam.SaveField //处理存储
	if !utils.IsValueEmpty(saveField) { //如果有保存字段，则处理
		template.AddParam(saveField, r.GetData())
	}
	// 处理判断服务是否运行正确
	tpl := handlerParam.TemplateTpl
	if tpl != nil { // 如果template 模板不为空，则渲染值
		success := utils.RenderTplBool(tpl, params)
		if !success { // 如果不成功，返回错误信息
			if handlerParam.ErrMsgTpl == nil {
				return common.NotOk(handlerParam.Template + "未配置，err_msg，无法渲染错误信息")
			} else { // 模板错误信息
				errMsg := utils.RenderTpl(handlerParam.ErrMsgTpl, params)
				return common.NotOk(errMsg)
			}
		}
	}
	return common.Ok(r.GetData(), "执行成功")
}

//todo template 存储结果，针对结果进行处理
func handlerParams(template *collect.Template, HandlerParams []collect.HandlerParam, ts *TemplateService) *common.Result {

	for _, handlerParam := range HandlerParams {
		ret := HandlerOneParams(&handlerParam, template, ts)
		if !ret.Success {
			return ret
		}
		//params := template.GetParams()
		//// 处理是否启用
		//if !utils.IsValueEmpty(handlerParam.Enable) { // 如果配置enable ，则判断是否启用
		//	enable := utils.RenderTplDataBool(handlerParam.EnableTpl, params)
		//	if !enable { // 如果结果为false 则跳过处理器
		//		continue
		//	}
		//
		//}
		//// 获取数据处理模块
		//module := GetModuleRegister(handlerParam.Key)
		//
		//if module == nil {
		//	return common.NotOk(handlerParam.Key + "处理模块没有找到")
		//}
		////数据处理
		//r := module.HandlerData(template, &handlerParam, ts)
		//if !r.Success { // 如果结果不成功直接返回
		//	return r
		//}
		////todo 这里处理count
		//saveField := handlerParam.SaveField //处理存储
		//if !utils.IsValueEmpty(saveField) { //如果有保存字段，则处理
		//	template.AddParam(saveField, r.GetData())
		//}
		//tpl := handlerParam.TemplateTpl
		//if tpl != nil { // 如果template 模板不为空，则渲染值
		//	success := utils.RenderTplBool(tpl, params)
		//	if !success { // 如果不成功，返回错误信息
		//		if handlerParam.ErrMsgTpl == nil {
		//			return common.NotOk(handlerParam.Template + "未配置，err_msg，无法渲染错误信息")
		//		} else { // 模板错误信息
		//			errMsg := utils.RenderTpl(handlerParam.ErrMsgTpl, params)
		//			return common.NotOk(errMsg)
		//		}
		//	}
		//}
	}
	return common.Ok(nil, "处理成功")
}

func (t *AfterLoader) ResultHandler(config collect.Plugin, template *collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {

	return handlerParams(template, template.ResultHandler, ts)
	//return common.Ok(nil, "参数处理完毕")
}
