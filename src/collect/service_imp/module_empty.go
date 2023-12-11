package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
)

type EmptyService struct {
	BaseHandler
}

func (s *EmptyService) Result(template *config.Template, ts *TemplateService) *common.Result {
	empty := make(map[string]interface{})
	return common.Ok(empty, "成功")
}
