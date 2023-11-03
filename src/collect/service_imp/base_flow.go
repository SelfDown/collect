package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"fmt"
	"time"
)

type BaseFlow struct {
}

const StartName = "start"
const EndName = "end"

/**
* 获取初始节点
 */
func getStartNode(dataDict map[string]*config.HandlerParam) *config.HandlerParam {
	startNode := dataDict[StartName]
	return startNode
}

/**
* 获取结束节点
 */
func getEndNode(dataDict map[string]*config.HandlerParam) *config.HandlerParam {
	startNode := dataDict[EndName]
	return startNode
}

func getServiceDict(serviceList []config.HandlerParam) (resultDict map[string]*config.HandlerParam, errMsg string) {
	d := make(map[string]*config.HandlerParam)
	for index, item := range serviceList {
		key := item.NodeKey
		if utils.IsValueEmpty(key) {
			return d, fmt.Sprintf("第【%d】个节点没有配置【node_key】属性", index+1)
		}
		nodeType := item.NodeType
		if utils.IsValueEmpty(nodeType) {
			return d, fmt.Sprintf("第【%d】个节点%s没有配置【node_type】属性", index+1, key)
		}
		name := item.Name
		if utils.IsValueEmpty(name) {
			return d, fmt.Sprintf("第【%d】个节点%s没有配置【name】属性", index+1, key)
		}
		// 除了尾节点，必须有next 属性
		nodeNext := item.NodeNext
		if nodeType != EndName && utils.IsValueEmpty(nodeNext) {
			return d, fmt.Sprintf("第【%d】个节点%s没有配置【node_next】属性", index+1, key)
		}
		if _, ok := d[key]; ok {
			// 如果已经注册过，则报错
			return d, fmt.Sprintf("第【%d】个节点已经存在【%s】节点，请检查配置", index+1, key)
		}
		d[key] = &serviceList[index]
	}
	startNode := getStartNode(d)
	if startNode == nil {
		return d, "流程未找到初始节点"
	}

	endNode := getEndNode(d)
	if endNode == nil {
		return d, "流程未找到结束节点"
	}
	return d, ""

}
func getNextNode(nodeStart *config.HandlerParam,
	nodeDict map[string]*config.HandlerParam,
	template *config.Template,
	nodeResult *common.Result) (*config.HandlerParam, string) {
	currentName := nodeStart.Name

	params := template.GetParams()
	// 如果有结果，先打印服务运行的结果
	if template.Log && nodeResult != nil {
		template.LogData(nodeResult)
	}
	// 如果没有运行结果，或则运行成功，找下个节点
	var nextNodeName string
	if nodeResult == nil || nodeResult.Success {
		next := nodeStart.NodeNextTpl
		nextNodeName = utils.RenderTplData(next, params).(string)
		if template.Log {
			template.LogData(nextNodeName)
		}
	} else {
		fail := nodeStart.NodeFailTpl
		if fail == nil {
			template.LogData("运行错误【" + currentName + "】但是未配置node_fail节点")
		}
		nextNodeName = utils.RenderTplData(fail, params).(string)
	}
	nextNode, ok := nodeDict[nextNodeName]
	if !ok {
		template.LogData("当前节点:")
		template.LogData(currentName)
		template.LogData("下个节点结果")
		template.LogData(nextNode)
		msg := fmt.Sprintf("流程结果返回错误，但是【%s】没有找到【%s】节点", currentName, nextNodeName)
		template.LogData(msg)
		return nil, msg
	}
	return nextNode, ""
}

/**
* 服务流程化
 */
func BaseHandlerNode(param *config.HandlerParam, template *config.Template, ts *TemplateService) *common.Result {
	// 必须拷贝处理一份，data_json 是指针
	tmp := utils.CopyWithPtr(param).(*config.HandlerParam)
	// 处理当个handler node
	return HandlerOneParams(tmp, template, ts)
}

/**
* 执行流程
 */
func (s *BaseFlow) _executeFlow(template *config.Template,
	ts *TemplateService,
	handlerNode func(param *config.HandlerParam, template *config.Template, ts *TemplateService) *common.Result) *common.Result {
	dataConfig := template.DataJsonConfig
	serviceList := dataConfig.Services
	if utils.IsValueEmpty(serviceList) {
		return common.NotOk("流程中services服务列表不能为空")
	}
	dict, errMsg := getServiceDict(serviceList)
	if !utils.IsValueEmpty(errMsg) {
		if template.Log {
			template.LogData(errMsg)
		}
		return common.NotOk(errMsg)
	}
	start := getStartNode(dict)
	last := start
	end := getEndNode(dict)
	current, errMsg := getNextNode(start, dict, template, nil)
	// 流程节点运行次数不能超过100
	count := 100
	i := 0
	firstErr := false
	var firstErrMsg string
	for {
		i += 1
		// 如果流程超过100次终止，防止流程进入死循环
		if i > count {
			errMsg = fmt.Sprintf("服务[%s]运行次数已经超过最大限制：%d", template.GetService(), count)
			return common.NotOk(errMsg)
		}
		// 如果是尾节点，则跳出循环
		if current == end {
			break
		}
		// 如果没有找到下个节点，则返回
		if current == nil {
			errMsg = fmt.Sprintf("服务[%s]运行到第[%d]个节点，找不到下个节点。上个节点[%s]", template.GetService(), i, last.NodeKey)
			return common.NotOk(errMsg)
		}
		// 计算程序耗时
		startClock := time.Now().Unix()
		// 运行服务结果
		nodeResult := handlerNode(current, template, ts)
		endClock := time.Now().Unix()
		if template.Log {
			msg := fmt.Sprintf("处理[%s]节点耗时%d秒", current.NodeKey, endClock-startClock)
			template.LogData(msg)
		}
		// 如果忽略错误运行，则将错误的结果改为True
		if current.IgnoreError && !nodeResult.Success {
			if template.Log {
				errMsg = fmt.Sprintf("忽略错误: 运行第【%d】节点【%s】失败 %s", i, current.Name, nodeResult.Msg)
				template.LogErr(errMsg)
			}
			nodeResult.Success = true
		}
		// 允许失败一次，出发fail 流程
		if !nodeResult.Success {
			if template.Log {
				errMsg = fmt.Sprintf("【%s】运行第【%d】个节点【%s】失败：%s", template.GetService(), i, current.Name, nodeResult.Msg)
				template.LogData(errMsg)
			}
			// 不能有2次流程错误
			if firstErr {
				errMsg = fmt.Sprintf("【%s】已经运行了2次错误，请检查流程节点", template.GetService())
				template.LogData(firstErrMsg)
				return common.NotOk(errMsg)
			}
			// 标记错误消息
			firstErr = true
			firstErrMsg = nodeResult.Msg

		}
		// 记录上次节点
		last = current
		// 流转到下个节点
		current, errMsg = getNextNode(current, dict, template, nodeResult)

	}
	// 如果有错误，就返回错误消息
	if firstErr {
		return common.NotOk(firstErrMsg)
	}
	empty := make(map[string]interface{})
	return common.Ok(empty, "成功")
}
func (s *BaseFlow) flow(template *config.Template, ts *TemplateService, handlerNode func(param *config.HandlerParam, template *config.Template, ts *TemplateService) *common.Result) *common.Result {
	// 执行中
	result := s._executeFlow(template, ts, handlerNode)
	// 处理finish
	return result
}
