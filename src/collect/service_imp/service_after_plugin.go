package collect

import (
	common "collect.mod/src/collect/common"
	"collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type AfterLoader struct {
}

func handlerParams(template collect.Template, HandlerParams []collect.HandlerParam, ts *TemplateService) *common.Result {

	for _, handlerParam := range HandlerParams {
		// 获取数据处理模块
		module := GetModuleRegister(handlerParam.Key)
		if module == nil {
			return common.NotOk(handlerParam.Key + "处理模块没有找到")
		}
		//数据处理
		r := module.HandlerData(&template, &handlerParam, ts)
		if !r.Success { // 如果结果不成功直接返回
			return r
		}
		//todo 这里处理count
		saveField := handlerParam.SaveField
		if !utils.IsValueEmpty(saveField) { //如果有保存字段，则处理
			template.AddParam(saveField, r.GetData())
		}
		tpl := handlerParam.TemplateTpl
		if tpl != nil { // 如果template 模板不为空，则渲染值
			success := utils.RenderTplBool(tpl, template.GetParams())
			if !success { // 如果不成功，返回错误信息
				if handlerParam.ErrMsgTpl == nil {
					return common.NotOk(handlerParam.Template + "未配置，err_msg，无法渲染错误信息")
				} else { // 模板错误信息
					errMsg := utils.RenderTpl(handlerParam.ErrMsgTpl, template.GetParams())
					return common.NotOk(errMsg)
				}
			}
		}
	}
	return common.Ok(nil, "处理成功")
}

func (t *AfterLoader) ResultHandler(config collect.Plugin, template collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {

	return handlerParams(template, template.ResultHandler, ts)
	//return common.Ok(nil, "参数处理完毕")
}
