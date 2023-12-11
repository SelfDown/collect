package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
	"net/url"
)

type File2Result struct {
	BaseHandler
}

func (uf *File2Result) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	ts.IsFileResponse = true
	ts.ResponseFilePath = utils.RenderVar(handlerParam.Path, params).(string)
	var filename string
	if utils.IsValueEmpty(handlerParam.ResultName) {
		filename = utils.FileName(ts.ResponseFilePath)
	} else {
		filename = utils.RenderVar(handlerParam.ResultName, params).(string)
	}

	ts.ResponseFileName = url.QueryEscape(filename)
	r := common.Ok(nil, "处理参数成功")
	return r
}
