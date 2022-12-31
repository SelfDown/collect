package simple

import (
	template_service "collect.mod/src/collect/service_imp"
	"fmt"
)

// UserCreate 注意只能在main中运行，否则配置文件路径不对
func Empty() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	//user := make(map[string]interface{})
	params := make(map[string]interface{})
	params["service"] = "hrm.empty_test"
	params["username"] = "zhangzhi"
	r := ts.Result(params, true)
	//r2 := ts.Result(params, true)
	fmt.Printf("%#v", r)
	//fmt.Printf("%#v", r2)
}
