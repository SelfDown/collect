package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"encoding/json"
)

type File2DataJson struct {
	BaseHandler
}

func (uf *File2DataJson) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	//params := template.GetParams()
	dataJsonValue := utils.RenderTpl(template.FileDataTpl, template.GetParams())
	var data map[string]interface{}
	json.Unmarshal([]byte(dataJsonValue), &data)
	r := common.Ok(data, "处理参数成功")
	return r
}
