package simple

import (
	"collect.mod/src/collect/template_service"
	"fmt"
)

// UserCreate 注意只能在main中运行，否则配置文件路径不对
func UserCreate() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	params["username"] = "zhangsan"
	params["nick"] = "张治"
	params["email"] = "1@163.com"
	params["statu"] = 1
	params["userpwd"] = "123456"
	params["address"] = "cs"
	params["tel"] = "18874948657"
	//params["page"] = 1
	params["service"] = "hrm.create_user"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
