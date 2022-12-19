package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	"fmt"
)

type Service2Field struct {
	BaseHandler
}

func (uf *Service2Field) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()

	params2 := make(map[string]interface{})
	params2["service"] = "hrm.empty_test2"
	r2 := ts.ResultInner(params2)
	fmt.Printf("%#v", r2)
	r := common.Ok(params, "处理参数成功")
	return r
}
