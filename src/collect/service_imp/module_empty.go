package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
)

type EmptyService struct {
	BaseHandler
}

func (s *EmptyService) Result(template *config.Template, ts *TemplateService) *common.Result {
	empty := make(map[string]interface{})
	return common.Ok(empty, "成功")
}
