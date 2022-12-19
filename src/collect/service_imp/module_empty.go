package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	"fmt"
)

type EmptyService struct {
	BaseHandler
}

func (s *EmptyService) Result(template *config.Template, ts *TemplateService) *common.Result {
	empty := make(map[string]interface{})
	fmt.Println(template.OpUser)
	return common.Ok(empty, "成功")
}
