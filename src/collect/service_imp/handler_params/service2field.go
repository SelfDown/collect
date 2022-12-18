package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	service_imp "collect.mod/src/collect/service_imp"
)

type Service2Field struct {
	service_imp.BaseHandler
}

func (uf *Service2Field) HandlerData(template *config.Template, handlerParam *config.HandlerParam) *common.Result {
	params := template.GetParams()
	//template.Result()

	//templateCopy := utils.Copy(*template).(config.Template)
	//template.Result()
	//params2 := make(map[string]interface{})
	//params2["service"] = "hrm.empty_test2"
	//templateCopy.SetParams(params2)
	//r2 := templateCopy.Result
	//fmt.Printf("%#v", r2)

	r := common.Ok(params, "处理参数成功")
	return r
}
