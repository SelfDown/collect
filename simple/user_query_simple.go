package simple

import (
	template_service "collect.mod/src/collect/service_imp"
	"fmt"
)

// QueryUserList QueryUserList注意只能在main中运行，否则配置文件路径不对
func QueryUserList() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	params["nick"] = "张治"
	params["name"] = "张治"
	//params["page"] = 1
	params["service"] = "hrm.user_list"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
