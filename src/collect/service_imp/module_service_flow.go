package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type ServiceFlowService struct {
	BaseHandler
	BaseFlow
}

// Result 服务流程化
func (s *ServiceFlowService) Result(template *config.Template, ts *TemplateService) *common.Result {
	// 流程返回结果
	flowResult := s.flow(template, ts, BaseHandlerNode)
	finish := template.DataJsonConfig.Finish
	if !utils.IsValueEmpty(finish.Key) {
		finishResult := HandlerOneParams(&finish, template, ts)
		if template.Log {
			template.LogData("finish运行结果")
			template.LogData(finishResult)
		}
	}

	return flowResult
}
