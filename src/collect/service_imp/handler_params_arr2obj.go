package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	"reflect"
)

type Arr2Obj struct {
	BaseHandler
}

func (ao *Arr2Obj) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	//params := template.GetParams()
	result := template.GetResult()
	if !template.HasResult() {
		return common.NotOk("没有找到结果")
	}
	data := result.Data
	dataArr := reflect.ValueOf(data)
	k := dataArr.Kind()
	if k == reflect.Slice || k == reflect.Array {
		if dataArr.Len() > 0 { // 如果之前有数据
			newResult := dataArr.Index(0).Interface()
			result.Data = newResult
		} else { // 如果没有数据，则返回空对象
			tmp := make(map[string]interface{}, 0)
			result.Data = tmp
		}
		template.SetResult(result)

	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
