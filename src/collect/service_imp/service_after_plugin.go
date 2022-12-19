package collect

import (
	common "collect.mod/src/collect/common"
	"collect.mod/src/collect/config"
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
		module.HandlerData(&template, &handlerParam, ts)
		// todo 这里处理save_field,template，err_msg 逻辑
	}
	return common.Ok(nil, "处理成功")
}

func (t *AfterLoader) ResultHandler(config collect.Plugin, template collect.Template, routerAll *collect.RouterAll, ts *TemplateService) *common.Result {

	return handlerParams(template, template.ResultHandler, ts)
	//return common.Ok(nil, "参数处理完毕")
}
