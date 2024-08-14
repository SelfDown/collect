package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"reflect"
)

type Arr2Obj struct {
	BaseHandler
}

func (ao *Arr2Obj) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	//params := template.GetParams()
	if utils.IsValueEmpty(handlerParam.Field) { // 如果没有field 则处理结果
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
	} else {
		params := template.GetParams()
		dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Field, params)
		if !utils.IsValueEmpty(errMsg) {
			return common.NotOk(errMsg)
		}
		varName := utils.GetRenderVarName(handlerParam.Field)
		params[varName] = dataList[0]
	}

	r := common.Ok(nil, "处理参数成功")
	return r
}
