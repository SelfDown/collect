package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"fmt"
	"strings"
)

type GetModifyData struct {
	BaseHandler
}
type BaseRule struct {
	Field    config.HandlerParam
	Template *config.Template
	Ts       *TemplateService
	HandlerData
}
type ChangeData struct {
	Before        interface{}
	BeforeDataMap map[string]interface{}
	After         interface{}
	AfterDataMap  map[string]interface{}
	Operation     string
}
type HandlerData interface {
	Handler() ([]map[string]interface{}, bool)
	Transfer(dataList []map[string]interface{}) []map[string]interface{}
}

func (s *BaseRule) Handler() {
	fmt.Println("base")
}
func (s *BaseRule) getChangeData(change *ChangeData) map[string]interface{} {
	data := make(map[string]interface{})
	// 设置改变前的值
	data[s.GetBeforeName()] = change.Before
	// 设置改变后的值
	data[s.GetAfterName()] = change.After
	// 设置操作
	op := s.Field.Operation
	if !utils.IsValueEmpty(change.Operation) {
		op = change.Operation
	}
	data[s.GetOperationName()] = op
	if s.Field.SaveOriginal {
		if op == removeOperation {
			data[s.GetValueName()] = change.Before
		} else if op == AddOperation {
			data[s.GetValueName()] = change.After
		}

	}

	// 设置名称
	data[s.GetName()] = s.Field.Name
	// 设置字段名称
	data[s.GetFieldName()] = s.Field.Field
	// 拼接右边对象的值
	rightFields := s.Field.AppendRightFields
	if len(rightFields) == 1 && rightFields[0] == "*" {

	}
	if rightFields != nil && len(rightFields) > 0 {
		for _, rightField := range rightFields {
			fieldName := utils.GetRenderVarName(rightField)
			value := utils.RenderVar(rightField, change.AfterDataMap)
			data[fieldName] = value
		}

	}
	return data
}
func (s *BaseRule) GetField() string {
	return s.Field.Field
}
func (s *BaseRule) GetLeftField() string {
	if !utils.IsValueEmpty(s.Field.LeftField) {
		return s.Field.LeftField
	}
	return s.GetField()
}
func (s *BaseRule) GetLeftValueField() string {

	return s.Field.LeftValueField
}
func (s *BaseRule) GetRightValueField() string {

	return s.Field.RightValueField
}
func (s *BaseRule) GetRightField() string {
	if !utils.IsValueEmpty(s.Field.RightField) {
		return s.Field.RightField
	}
	return s.GetField()
}
func (s *BaseRule) GetBeforeName() string {
	return "before"
}
func (s *BaseRule) GetName() string {
	return "name"
}
func (s *BaseRule) GetFieldName() string {
	return "field"
}
func (s *BaseRule) GetAfterName() string {
	return "after"
}
func (s *BaseRule) GetOperationName() string {
	return "operation"
}
func (s *BaseRule) GetValueName() string {
	return "value"
}
func (s *BaseRule) LeftObj() map[string]interface{} {
	params := s.Template.GetParams()
	if !utils.IsValueEmpty(s.Field.Left) {
		subParams := utils.RenderVar(s.Field.Left, params)
		return subParams.(map[string]interface{})
	}
	return params
}

func (s *BaseRule) LeftObjArr() []map[string]interface{} {
	params := s.Template.GetParams()
	dataList := make([]map[string]interface{}, 0)
	if !utils.IsValueEmpty(s.Field.Left) {
		subParams, _ := utils.RenderVarToArrMap(s.Field.Left, params)

		return subParams
	}
	return dataList
}
func (s *BaseRule) RightObj() map[string]interface{} {
	params := s.Template.GetParams()
	if !utils.IsValueEmpty(s.Field.Right) {
		subParams := utils.RenderVar(s.Field.Right, params)
		return subParams.(map[string]interface{})
	}
	return params
}

func (s *BaseRule) RightObjArr() []map[string]interface{} {
	params := s.Template.GetParams()
	dataList := make([]map[string]interface{}, 0)
	if !utils.IsValueEmpty(s.Field.Right) {
		subParams, _ := utils.RenderVarToArrMap(s.Field.Right, params)

		return subParams
	}
	return dataList
}

func (s *BaseRule) Transfer(dataList []map[string]interface{}) []map[string]interface{} {
	// 如果没有配置service，直接返回
	if utils.IsValueEmpty(s.Field.Service) {
		return dataList
	}
	//获取当前值的列表
	currentValueList := make([]interface{}, 0)
	for _, item := range dataList {
		before := item[s.GetBeforeName()]
		after := item[s.GetAfterName()]
		currentValueList = append(currentValueList, before, after)
	}
	// 转换数据,调用服务
	params := s.Template.GetParams()
	valueKey := s.Field.ValueListField
	oldValue, hasKey := params[valueKey]
	// 将值列表设置进去，以便查找
	params[valueKey] = currentValueList
	// 拼装服务
	serviceParam := utils.GetServiceParam(s.Field.Service, params, s.Field.AppendParam)
	r2 := s.Ts.ResultInner(serviceParam)
	if !r2.Success {
		return dataList
	}
	if hasKey { // 如果存在就还原
		params[valueKey] = oldValue
	} else { // 不存在就删除
		delete(params, valueKey)
	}
	tData := r2.GetData().([]map[string]interface{})
	transDict := make(map[string]string)
	// 将数据转成字典
	key := s.Field.TargetTransferKey
	value := s.Field.TargetTransferValue
	for _, item := range tData {
		k := utils.RenderVar(key, item).(string)
		v := utils.RenderVar(value, item).(string)
		transDict[k] = v
	}

	for index, item := range dataList {
		before := item[s.GetBeforeName()].(string)
		if !utils.IsValueEmpty(before) {
			dataList[index][s.GetBeforeName()] = transferValue(before, transDict)
		}
		after := item[s.GetAfterName()].(string)
		if !utils.IsValueEmpty(after) {
			dataList[index][s.GetAfterName()] = transferValue(after, transDict)
		}
	}
	return dataList

}
func transferValue(original string, transDict map[string]string) string {
	value, ok := transDict[original]
	if !ok {
		return original
	}
	return value
}

func (uf *GetModifyData) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	changData := make([]map[string]interface{}, 0)
	for _, field := range template.ModifyConfigData.Fields {
		rule := field.Rule
		baseRule := BaseRule{
			Field:    field,
			Template: template,
			Ts:       ts,
		}
		var fieldRule HandlerData
		// 如果是简单字段对比
		if SimpleFieldRuleName == rule {
			fieldRule = &SimpleFieldRule{
				BaseRule: baseRule,
			}
		} else if SimpleArrayRuleName == rule {
			fieldRule = &SimpleArrayRule{
				BaseRule: baseRule,
			}
		} else if ArrayObjRuleName == rule {
			fieldRule = &ArrayObjRule{
				BaseRule: baseRule,
			}
		}
		changeList, hasChange := fieldRule.Handler()
		if hasChange {
			// 转换数据
			fieldRule.Transfer(changeList)
			changData = append(changData, changeList...)
		}

	}
	r := common.Ok(changData, "处理参数成功")
	return r
}

const ModifyOperation = "modify"
const AddOperation = "add"
const removeOperation = "remove"

// SimpleFieldRuleName 简单的字段比对
const SimpleFieldRuleName = "compare_field_value"

type SimpleFieldRule struct {
	BaseRule
}

func (s *SimpleFieldRule) Handler() ([]map[string]interface{}, bool) {
	leftData := s.LeftObj()
	rightData := s.RightObj()
	leftValue := utils.RenderVar(s.GetLeftField(), leftData)
	rightValue := utils.RenderVar(s.GetRightField(), rightData)
	dataList := make([]map[string]interface{}, 0)
	if leftValue == rightValue {
		return dataList, false
	}
	change := ChangeData{
		AfterDataMap:  rightData,
		BeforeDataMap: leftData,
		Before:        leftValue,
		After:         rightValue,
	}
	data := s.getChangeData(&change)
	dataList = append(dataList, data)

	return dataList, true
}

// SimpleArrayRuleName 简单的数组对比，以逗号分割
const SimpleArrayRuleName = "simple_array_value"

type SimpleArrayRule struct {
	BaseRule
}

func string2Arr(value interface{}) []string {
	if utils.IsValueEmpty(value) {
		return make([]string, 0)
	}
	arr := strings.Split(value.(string), ",")
	return arr
}
func getArrayMap(data []string) map[string]int {
	dict := make(map[string]int)
	for _, item := range data {
		dict[item] = 1
	}
	return dict
}
func getNotExistsData(data []string, dict map[string]int) []string {
	dataList := make([]string, 0)
	for _, item := range data {
		if _, ok := dict[item]; !ok {
			dataList = append(dataList, item)
		}
	}
	return dataList
}
func (s *SimpleArrayRule) Handler() ([]map[string]interface{}, bool) {
	leftData := s.LeftObj()
	rightData := s.RightObj()
	leftValue := utils.RenderVar(s.GetLeftField(), leftData)
	rightValue := utils.RenderVar(s.GetRightField(), rightData)

	//将字符串转出字典，获取差集
	leftArr := string2Arr(leftValue)
	leftDict := getArrayMap(leftArr)
	rightArr := string2Arr(rightValue)
	rightDict := getArrayMap(rightArr)
	// 新增的数据
	addList := getNotExistsData(leftArr, rightDict)
	// 删除的数据
	removeList := getNotExistsData(rightArr, leftDict)
	dataList := make([]map[string]interface{}, 0)
	if len(addList) <= 0 && len(removeList) <= 0 {
		return dataList, false
	}
	//处理新增
	for _, item := range addList {
		change := ChangeData{
			AfterDataMap:  rightData,
			BeforeDataMap: leftData,
			Before:        "",
			After:         item,
			Operation:     AddOperation,
		}
		data := s.getChangeData(&change)
		dataList = append(dataList, data)
	}
	//处理删除
	for _, item := range removeList {
		change := ChangeData{
			AfterDataMap:  rightData,
			BeforeDataMap: leftData,
			Before:        item,
			After:         "",
			Operation:     removeOperation,
		}
		data := s.getChangeData(&change)
		dataList = append(dataList, data)
	}
	return dataList, true
}

// ArrayObjRuleName 简单的数组对比，以逗号分割
const ArrayObjRuleName = "array_obj_value"

type ArrayObjRule struct {
	BaseRule
}

func getArrayObjMap(data []map[string]interface{}, fieldName string) map[string]int {
	dict := make(map[string]int)
	for _, item := range data {
		key := utils.RenderVar(fieldName, item).(string)
		dict[key] = 1
	}
	return dict
}
func getArrNotExistsData(data []map[string]interface{}, dict map[string]int, fieldName string) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0)
	for _, item := range data {
		key := utils.RenderVar(fieldName, item).(string)
		if _, ok := dict[key]; !ok {
			dataList = append(dataList, item)
		}
	}
	return dataList
}

func (s *ArrayObjRule) handlerAddRemove(leftArr []map[string]interface{}, rightArr []map[string]interface{}, leftDict map[string]int, rightDict map[string]int) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0)
	addList := getArrNotExistsData(leftArr, rightDict, s.GetLeftField())
	removeList := getArrNotExistsData(rightArr, leftDict, s.GetRightField())
	//todo 比较key相同，然后其他字段不同

	if len(addList) <= 0 && len(removeList) <= 0 {
		return dataList
	}
	none := make(map[string]interface{})
	//处理新增
	for _, item := range addList {
		after := utils.RenderVar(s.GetLeftField(), item)
		change := ChangeData{
			AfterDataMap:  item,
			BeforeDataMap: none,
			After:         after,
			Before:        "",
			Operation:     AddOperation,
		}
		data := s.getChangeData(&change)
		dataList = append(dataList, data)
	}
	//处理删除
	for _, item := range removeList {
		before := utils.RenderVar(s.GetRightField(), item)
		change := ChangeData{
			AfterDataMap:  item,
			BeforeDataMap: none,
			After:         "",
			Before:        before,
			Operation:     removeOperation,
		}
		data := s.getChangeData(&change)
		dataList = append(dataList, data)
	}
	return dataList

}

func (s *ArrayObjRule) handlerModify(leftArr []map[string]interface{}, rightArr []map[string]interface{}, leftDict map[string]int, rightDict map[string]int) []map[string]interface{} {
	dataList := make([]map[string]interface{}, 0)

	// 获取左边的字典+数组
	leftCommon := make([]map[string]interface{}, 0)
	leftCommonDict := make(map[string]map[string]interface{})
	for _, item := range leftArr {
		key := utils.RenderVar(s.GetLeftField(), item).(string)
		if _, ok := rightDict[key]; ok {
			leftCommon = append(leftCommon, item)
			leftCommonDict[key] = item
		}
	}
	// 获取右边的字典+数组
	rightCommon := make([]map[string]interface{}, 0)
	rightCommonDict := make(map[string]map[string]interface{})
	for _, item := range rightArr {
		key := utils.RenderVar(s.GetRightField(), item).(string)
		if _, ok := leftDict[key]; ok {
			rightCommon = append(rightCommon, item)
			rightCommonDict[key] = item
		}
	}

	for _, leftObj := range leftCommon {
		key := utils.RenderVar(s.GetLeftField(), leftObj).(string)
		rightObj := rightCommonDict[key]
		leftValue := utils.RenderVar(s.GetLeftValueField(), leftObj)
		rightValue := utils.RenderVar(s.GetRightValueField(), rightObj)
		if leftValue == rightValue {
			continue
		}
		//before := utils.RenderVar(s.GetRightField(), item)
		change := ChangeData{
			AfterDataMap:  rightObj,
			BeforeDataMap: leftObj,
			After:         leftValue,
			Before:        rightValue,
			Operation:     ModifyOperation,
		}
		data := s.getChangeData(&change)
		dataList = append(dataList, data)
	}

	return dataList

}

func (s *ArrayObjRule) Handler() ([]map[string]interface{}, bool) {
	dataList := make([]map[string]interface{}, 0)
	leftArr := s.LeftObjArr()
	leftDict := getArrayObjMap(leftArr, s.GetLeftField())
	rightArr := s.RightObjArr()
	rightDict := getArrayObjMap(rightArr, s.GetRightField())
	if s.Field.WithAddRemove {
		list := s.handlerAddRemove(leftArr, rightArr, leftDict, rightDict)
		dataList = append(dataList, list...)
	}
	if !utils.IsValueEmpty(s.Field.RightValueField) {
		list := s.handlerModify(leftArr, rightArr, leftDict, rightDict)
		dataList = append(dataList, list...)
	}

	return dataList, true
}
