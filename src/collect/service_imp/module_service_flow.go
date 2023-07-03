package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
)

type ServiceFlowService struct {
	BaseHandler
	BaseFlow
}

///**
//* 执行流程
// */
//func (s *ServiceFlowService) _executeFlow(template *config.Template, ts *TemplateService) *common.Result {
//	dataConfig := template.DataJsonConfig
//	serviceList := dataConfig.Services
//	if utils.IsValueEmpty(serviceList) {
//		return common.NotOk("流程中services服务列表不能为空")
//	}
//	empty := make(map[string]interface{})
//	return common.Ok(empty, "成功")
//}
//func (s *ServiceFlowService) flow(template *config.Template, ts *TemplateService) *common.Result {
//	s._executeFlow(template, ts)
//	return common.Ok(nil, "成功")
//}

func (s *ServiceFlowService) Result(template *config.Template, ts *TemplateService) *common.Result {
	// 流程返回结果
	flowResult := s.flow(template, ts, BaseHandlerNode)
	return flowResult
}
