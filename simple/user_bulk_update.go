package simple

import (
	template_service "collect.mod/src/collect/service_imp"
	"fmt"
)

// UserCreate 注意只能在main中运行，否则配置文件路径不对
func UserUpdateBulk() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	user := make(map[string]interface{})

	user["username"] = "zhangsan"
	user["nick"] = "张治"
	user["userid"] = "z1"
	user["email"] = "1@163.com"
	user["statu"] = 0
	user["userpwd"] = "123456"
	user["address"] = "cs11"
	user["tel"] = "18874948657"

	user2 := make(map[string]interface{})
	user2["userid"] = "a1"
	user2["username"] = "zhangsan"
	user2["nick"] = "张治1"
	user2["email"] = "1@163.com"
	user2["statu"] = 1
	user2["userpwd"] = "123456"
	user2["address"] = "cs"
	user2["tel"] = "18874948657"
	//params["page"] = 1
	params := make(map[string]interface{})
	userList := make([]map[string]interface{}, 0)
	userList = append(userList, user)
	userList = append(userList, user2)
	//user_list :=[]map[string][]
	params["service"] = "hrm.bulk_update_user"
	params["user_list"] = userList
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
}
