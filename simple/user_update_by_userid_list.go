package simple

import (
	"collect.mod/src/collect/template_service"
	"fmt"
)

// UserUpdateByUsernameNick 注意只能在main中运行，否则配置文件路径不对
func UserUpdateByUseridList() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	params["userid_list"] = [...]string{"3c571ff1-ead0-4e3e-837f-2a2f03cf1ee6",
		"5c8eb14c-c9d5-4f04-a15f-0ba9df90567a",
		"25668da7-ebfc-4875-ab0c-7fa333957034"}
	params["nick"] = "张治1"
	params["username"] = "zhangsan"
	params["email"] = "1@163.com"
	params["statu"] = 0
	params["userpwd"] = "123456"
	params["address"] = "长沙1"
	params["tel"] = "18874948657"
	//params["page"] = 1
	params["service"] = "hrm.update_user_by_userid_list"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
