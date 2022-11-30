package main

import (
	"fmt"

	config "collect.mod/src/collect/config"
	startup "collect.mod/src/collect/startup"

	template_service "collect.mod/src/collect/template_service"
	// "plugin"

	utils "collect.mod/src/collect/utils"
)

func init() {
	// 加载配置
	utils.LoadAppProperties("./conf/application.properties")
	t := config.Template{}
	// 加载系统插件，主要加载count、file_data,
	routerAll := startup.LoadSystemServices(&t)
	//获取启动注册的服务列表，然后路由设置注册服务
	routerAll.SetRegisterList(startup.GetRegisterList())
	//设置服务
	config.SetLocalRouter(routerAll)

}

func main() {
	ts := template_service.TemplateService{OpUser: "zhangzhi"}
	params := make(map[string]interface{})
	params["nick"] = "张治"
	//params["page"] = 1
	params["service"] = "hrm.user_list"
	r := ts.Result(params, true)
	fmt.Printf("%#v", r)
	// params = make(map[string]interface{})
	// params["nick"] = "兰雄"
	// params["status"] = 1
	// params["service"] = "hrm.user_list"
	// r = ts.Result(params, true)
	// fmt.Println(r.GetData())
	// fmt.Println(r.GetMsg())
	// fmt.Println(r.GetSuccess())

	// fmt.Printf("ts: %v\n", ts)

}
