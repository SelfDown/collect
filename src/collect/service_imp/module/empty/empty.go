package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	serviceImp "collect.mod/src/collect/service_imp"
	"fmt"
)

type EmptyService struct {
	serviceImp.BaseHandler
}

func (s *EmptyService) Result(template *config.Template) *common.Result {
	empty := make(map[string]interface{})
	fmt.Println(template.OpUser)
	return common.Ok(empty, "成功")
}
