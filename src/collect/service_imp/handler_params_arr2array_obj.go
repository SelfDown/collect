package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"reflect"
)

type Arr2arrayObj struct {
	BaseHandler
}

func (uf *Arr2arrayObj) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	foreach := utils.RenderVar(handlerParam.Foreach, params)
	if !utils.IsArray(foreach) {
		return common.NotOk(handlerParam.Foreach + "不是数组")
	}
	if utils.IsValueEmpty(handlerParam.Item) {
		return common.NotOk("foreach未配置item")
	}
	// 先备份一波，如果item 的key存在
	itemKey := handlerParam.Item
	oldValue := params[itemKey]
	IsItemKeyEmpty := utils.IsEmpty(itemKey, params)
	fields := handlerParam.Fields
	re := reflect.ValueOf(foreach)
	dataList := make([]map[string]interface{}, re.Len())
	// 数组循环进行赋值
	for i := 0; i < re.Len(); i++ {
		value := re.Index(i).Interface()
		//改了整体参数item 的key，最后还原
		params[itemKey] = value
		dataItem := make(map[string]interface{})
		// 挨个字段进行赋值
		for _, field := range fields {
			dataItem[field.Field] = utils.RenderTplData(field.TemplateTpl, params)
		}
		dataList[i] = dataItem
	}

	// 还原item 的key
	if !IsItemKeyEmpty {
		params[itemKey] = oldValue
	} else { // 不存在就删除
		delete(params, itemKey)
	}
	r := common.Ok(dataList, "处理参数成功")
	return r
}
