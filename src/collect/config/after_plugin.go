package collect

import common "collect.mod/src/collect/common"

type AfterLoader struct {
}

func handlerParams(template Template, HandlerParams []HandlerParam) *common.Result {
	for _, handlerParam := range HandlerParams {
		// 获取数据处理模块
		module := GetModuleRegister(handlerParam.Key)
		if module == nil {
			return common.NotOk(handlerParam.Key + "处理模块没有找到")
		}
		//数据处理
		module.HandlerData(&template, &handlerParam)
	}
	return common.Ok(nil, "处理成功")
}

func (t *AfterLoader) ResultHandler(config Plugin, template Template, routerAll *RouterAll) *common.Result {

	return handlerParams(template, template.ResultHandler)
	//return common.Ok(nil, "参数处理完毕")
}
