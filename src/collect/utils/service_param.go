package collect

func GetServiceParam(serviceParam map[string]interface{}, params map[string]interface{}, appendParam bool) map[string]interface{} {
    serviceCopy :=Copy(serviceParam).(map[string]interface{})
	for key, value := range serviceCopy {
		//不允许从参数中直接调用service，以防请求参数乱改，恶意调用其他服务不安全
		if key == "service" {
			continue
		}
		valueStr, ok := value.(string)
		// 判断是否为，参数变量，如果参数变量直接取参数值
		if ok && IsRenderVar(valueStr) {
			val := RenderVar(valueStr, params)
			serviceCopy[key] = val
		}
	}
	//拼接剩余参数
	if appendParam {
		for key, value := range params {
			if _, ok := serviceCopy[key]; !ok {
				serviceCopy[key] = value
			}
		}
	}
	return serviceCopy

}
