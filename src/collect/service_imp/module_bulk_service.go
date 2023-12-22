package collect

import (
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"math"
)

type BulkService struct {
	BaseHandler
}

func runService(params map[string]interface{}, ts *TemplateService, result chan<- *common.Result) {

	r2 := ts.ResultInner(params)
	result <- r2
}

func runServiceList(foreach []map[string]interface{}, params map[string]interface{}, batch config.HandlerParam, ts *TemplateService, serviceMap map[string]interface{}) []map[string]interface{} {
	//foreach := utils.RenderVar(batch.Foreach, params).([]map[string]interface{})
	resultChanList := make([]chan *common.Result, len(foreach))
	resultList := make([]map[string]interface{}, len(foreach))
	transfer := utils.IsValueEmpty(serviceMap)
	// 生成参数
	cpyParams := utils.Copy(params)
	for index, item := range foreach {
		batchService := utils.Copy(batch.Service).(map[string]interface{})
		//设置params
		item["params"] = cpyParams
		serviceParam := utils.GetServiceParam(batchService, item, batch.AppendItemParam)
		//存在转换的map就转换
		if !transfer {
			serviceName := "service"
			// 获取原有的旧服务
			oldName := serviceParam[serviceName].(string)
			oldService := utils.RenderVar(oldName, item).(string)
			// 将原有的服务转
			targetService := serviceMap[oldService]
			serviceParam[serviceName] = targetService
		}

		ch := make(chan *common.Result)
		resultChanList[index] = ch
		go runService(serviceParam, ts, resultChanList[index])
		resultList[index] = serviceParam
	}
	// 获取结果对象
	for index, result := range resultChanList {
		resultData := <-result
		resultList[index][batch.SaveField] = resultData
	}
	return resultList

}

func (s *BulkService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	batch := template.Batch
	foreach, ok := utils.RenderVarToArrMap(batch.Foreach, params)
	if !utils.IsValueEmpty(ok) {
		return common.NotOk(ok)
	}
	fileData := template.FileData
	serviceMap := make(map[string]interface{})
	if !utils.IsValueEmpty(fileData) {
		json.Unmarshal([]byte(fileData), &serviceMap)
	}

	//resultChanList := make([]chan *common.Result, len(foreach))
	l := len(foreach)
	resultList := make([]map[string]interface{}, 0)
	size := 50
	times := gocast.ToInt(math.Ceil(float64(len(foreach)) / float64(size)))

	for i := 0; i < times; i++ {
		start := i * size
		end := (i + 1) * size
		if end > l {
			end = l
		}
		forItem := foreach[start:end]
		itemResultList := runServiceList(forItem, params, batch, ts, serviceMap)
		resultList = append(resultList, itemResultList...)
	}

	return common.Ok(resultList, "成功")
}
