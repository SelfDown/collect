package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
)

type BaseHandler struct {
}

func (s *BaseHandler) Result(template *config.Template) *common.Result {
	return common.Ok(nil, "")
}

func (s *BaseHandler) HandlerData(template *config.Template, handlerParam *config.HandlerParam) *common.Result {
	return common.Ok(nil, "")
}
