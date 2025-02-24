package collect

//GetRegisterList  获取注册列表，对象名称绑定key，一定要注意对象名称一定要唯一

func GetRegisterList() []ModuleResult {
	l := make([]ModuleResult, 0)
	// handler_params 参数处理
	l = append(l, &UpdateField{})
	l = append(l, &CheckField{})
	l = append(l, &UpdateArray{})
	l = append(l, &Service2Field{})
	l = append(l, &Arr2Obj{})
	l = append(l, &Param2Result{})
	l = append(l, &Param2Count{})
	l = append(l, &Params2Result{})
	l = append(l, &SessionAdd{})
	l = append(l, &SessionGet{})
	l = append(l, &SessionRemove{})
	l = append(l, &Data2Excel{})
	l = append(l, &File2Result{})
	l = append(l, &Excel2Data{})
	l = append(l, &IgnoreData{})
	l = append(l, &Result2Params{})
	l = append(l, &Result2Map{})
	l = append(l, &Count2Map{})
	l = append(l, &File2DataJson{})
	l = append(l, &Field2Array{})
	l = append(l, &Arr2arrayObj{})
	l = append(l, &GetModifyData{})
	l = append(l, &FilterArr{})
	l = append(l, &PropArr{})
	l = append(l, &Arr2Dict{})
	l = append(l, &HandlerCache{})
	l = append(l, &GroupBy{})
	l = append(l, &UpdateArrayFromArray{})
	l = append(l, &CombineArray{})
	l = append(l, &PreventDuplication{})
	l = append(l, &ToTree{})
	l = append(l, &UpdateOrder{})
	l = append(l, &ToList{})
	l = append(l, &OrderBy{})
	l = append(l, &Agg{})
	// module 模块处理
	l = append(l, &ModelSaveService{})
	l = append(l, &ModelUpdateService{})
	l = append(l, &ModelDeleteService{})
	l = append(l, &BulkCreateService{})
	l = append(l, &BulkUpsertService{})
	l = append(l, &SqlService{})
	l = append(l, &EmptyService{})
	l = append(l, &HttpService{})
	l = append(l, &LdapService{})
	l = append(l, &ServiceFlowService{})
	l = append(l, &BulkService{})

	return l
}
