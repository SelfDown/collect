package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type UpdateArrayFromArray struct {
	BaseHandler
}

func (uf *UpdateArrayFromArray) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	// 直接渲染变量
	dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	// 直接渲染变量
	rightList, errMsg := utils.RenderVarToArrMap(handlerParam.Right, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	rightDict := make(map[string]map[string]interface{})
	leftFieldName := handlerParam.Field
	rightFieldName := handlerParam.Field
	if !utils.IsValueEmpty(handlerParam.RightField) {
		rightFieldName = handlerParam.RightField
	}
	for _, rightItem := range rightList {
		rightField := utils.RenderVar(rightFieldName, rightItem).(string)
		rightDict[rightField] = rightItem
	}
	//如果是空，填充右边的字段，处理左边字段不对等
	if handlerParam.NoneFillRight {
		nameDict := make(map[string]string)
		nameList := make([]string, 0)
		//遍历数组拿出所有key，如果是有性能问题可以只跳前面100个
		for _, item := range dataList {
			for key, _ := range item {
				if _, ok := nameDict[key]; !ok {
					nameDict[key] = key
					nameList = append(nameList, key)
				}
			}
		}
		for _, item := range dataList {
			leftField := utils.RenderVar(leftFieldName, item).(string)
			rightItem, ok := rightDict[leftField]
			if !ok {
				continue
			}
			//取出左边所有独立的key
			for _, key := range nameList {
				if _, itemOk := item[key]; !itemOk {
					item[key] = rightItem[key]
				}
			}
		}

	}

	// 如果存在fields，todo 暂时未实现
	if len(handlerParam.Fields) > 0 {
		var paramsCopy map[string]interface{}
		if !utils.IsValueEmpty(handlerParam.Item) { // 如果没有配置item 则取本身
			paramsCopy = utils.CopyMap(params)
		}
		for _, field := range handlerParam.Fields {

			for _, item := range dataList {
				if !utils.IsValueEmpty(handlerParam.Item) { // 如果配置了item，设置item
					paramsCopy[utils.ItemName] = item
				} else { // 没有配置item取整个item
					paramsCopy = item
				}
				if !utils.IsValueEmpty(handlerParam.Right) {
					fieldName := handlerParam.Field
					fieldId := utils.RenderVar(fieldName, item).(string)
					rightItem, ok := rightDict[fieldId]
					if ok {
						paramsCopy[utils.RightName] = rightItem
					} else { // todo 这里不确定没有匹配上是要跳过，还是填充key，先跳过
						continue
					}

				}
				// 处理判断不需要修改的，比如它本来就传了值
				if !utils.IsValueEmpty(handlerParam.IfTemplate) {
					run := utils.RenderTplBool(handlerParam.IfTemplateTpl, paramsCopy)
					if !run {
						continue
					}
				}

				//渲染值
				value := utils.RenderTplDataWithType(field.TemplateTpl, paramsCopy, field.Type)
				item[field.Field] = value
			}
		}
	}

	if template.Log {
		template.LogData("update_array 处理结果")
		template.LogData(params)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
