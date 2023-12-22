package collect

import (
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type File2DataJson struct {
	BaseHandler
}

func (uf *File2DataJson) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	//params := template.GetParams()
	dataJsonValue := utils.RenderTpl(template.FileDataTpl, template.GetParams())
	var data map[string]interface{}
	json.Unmarshal([]byte(dataJsonValue), &data)
	if data == nil {
		var dataArr []map[string]interface{}
		json.Unmarshal([]byte(dataJsonValue), &dataArr)
		return common.Ok(dataArr, "参数处理成功")
	}
	r := common.Ok(data, "处理参数成功")
	return r
}
