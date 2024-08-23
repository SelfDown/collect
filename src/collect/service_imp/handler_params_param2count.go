package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
)

type Param2Count struct {
	BaseHandler
}

func (pr *Param2Count) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	result := utils.RenderVar(handlerParam.Field, params)
	count := gocast.ToInt64(result)
	template.SetCount(count)
	template.GetResult().Count = count
	r := common.Ok(nil, "处理参数成功")
	return r
}
