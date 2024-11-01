package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
)

type Agg struct {
	BaseHandler
}

func (uf *Agg) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	fields := handlerParam.Fields
	var paramsCopy map[string]interface{}
	if !utils.IsValueEmpty(handlerParam.Item) { // 如果没有配置item 则取本身
		paramsCopy = utils.CopyMap(params)
	}
	for _, item := range arr {
		fieldMap := make(map[string]interface{})
		// 初始化字段信息
		for _, field := range fields {
			// 目标字段
			fieldName := utils.GetRenderVarName(field.To)
			fieldMap[fieldName] = 0
		}
		for _, field := range fields {

			fieldToName := utils.GetRenderVarName(field.To)
			fieldFromName := field.From

			childrenList, errMsg := utils.RenderVarToArrMap(childName, item)
			if !utils.IsValueEmpty(errMsg) {
				continue
			}
			// 计算统计值
			for _, child := range childrenList {
				if !utils.IsValueEmpty(field.Template) { // 如果template 不为空，计算是否统计
					paramsCopy[utils.ItemName] = child
					// 如果验证不通过，则不统计
					ok := utils.RenderTplBool(field.TemplateTpl, paramsCopy)
					if !ok {
						continue
					}
				}
				// 获取当前值
				value := utils.RenderVar(fieldFromName, child)
				if field.Type == "count" {
					fieldMap[fieldToName] = gocast.ToInt64(fieldMap[fieldToName]) + 1
				} else if field.Type == "sum" {
					fieldMap[fieldToName] = gocast.ToFloat64(fieldMap[fieldToName]) + gocast.ToFloat64(value)
				} else if field.Type == "avg" {

				} else if field.Type == "max" {

				} else if field.Type == "min" {

				} else if field.Type == "group_concat" {

				}
			}

		}
		// 添加字段信息
		for key, value := range fieldMap {
			item[key] = value
		}

	}

	r := common.Ok(arr, "处理参数成功")
	return r
}
