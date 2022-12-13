package simple

import (
	"collect.mod/src/collect/template_service"
	"fmt"
)

// UserUpdateAll 如果用户ID为空，则更新
func UserUpdateAll() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	params["address"] = "长沙1"
	params["service"] = "hrm.update_user_all"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
