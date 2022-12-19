package simple

import (
	template_service "collect.mod/src/collect/service_imp"
	"fmt"
)

// UserUpdateByUsernameNick 注意只能在main中运行，否则配置文件路径不对
func UserUpdateByUsernameNick() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	//params["userid"] = "25668da7-ebfc-4875-ab0c-7fa333957034"
	params["nick"] = "张治1"
	params["username"] = "zhangsan"
	params["email"] = "1@163.com"
	params["statu"] = 0
	params["userpwd"] = "123456"
	params["address"] = "长沙"
	params["tel"] = "18874948657"
	//params["page"] = 1
	params["service"] = "hrm.update_user_by_username_nick"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
