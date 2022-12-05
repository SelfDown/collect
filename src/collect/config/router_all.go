package collect

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"

	text_template "text/template"

	utils "collect.mod/src/collect/utils"
)

// type ModuleResult interface {
// 	Result(template *template.Template) (*common.Result, error)
// }
/**
* 项目总路由
 */
type RouterAll struct {
	Services           []RouterProject                      // 项目路由
	CurrentDir         string                               // 当前路径目录
	Path               string                               // 当前路径
	serviceDict        map[string]map[string]*ServiceConfig // 服务的二维字典
	moduleRegisterDict map[string]ModuleResult              //服务注册对象
	registerList       []ModuleResult                       // 设置注册列表
	LoadStartupPlugin  []Plugin                             `yaml:"load_startup_plugin"` // 加载启动插件
	FileContentPlugin  []Plugin                             `yaml:"file_content_plugin"` // 文件内容处理插件
	ModuleHandler      []Plugin                             `yaml:"module_handler"`      // 文件内容处理插件
	DataHandler        []Plugin                             `yaml:"data_handler"`        // 文件内容处理插件
	BeforePlugin       []Plugin                             `yaml:"before_plugin"`       // 文件内容处理插件
	AfterPlugin        []Plugin                             `yaml:"after_plugin"`        // 文件内容处理插件
	moduleDict         map[string]Plugin                    //服务注册对象字典
}

func (t *RouterAll) SetRegisterList(registerList []ModuleResult) {
	// 设置转换字典
	moduleDict := make(map[string]Plugin)
	for _, module := range t.ModuleHandler {
		moduleDict[module.Path] = module
	}
	for _, module := range t.DataHandler {
		moduleDict[module.Path] = module
	}
	t.moduleDict = moduleDict
	// 清空字典
	t.moduleRegisterDict = make(map[string]ModuleResult)
	for _, reg := range registerList {
		// 这里根据字符串，注册层服务
		name := reflect.TypeOf(reg).Elem().Name()
		//moduleRegisterDict 进行赋值
		if module, ok := moduleDict[name]; ok {
			t.moduleRegisterDict[module.Key] = reg
		} else {
			log.Println("模块【" + name + "】没有注册，请检查配置！！！")
		}

	}

	t.registerList = registerList
}
func (t *RouterAll) GetRegisterList() []ModuleResult {
	return t.registerList
}

func (t *RouterAll) GetModuleRegister(name string) ModuleResult {
	return t.moduleRegisterDict[name]
}
func (t *RouterAll) SetModuleRegister(name string, module ModuleResult) {
	if t.moduleRegisterDict == nil {
		t.moduleRegisterDict = make(map[string]ModuleResult, 0)
	}
	t.moduleRegisterDict[name] = module
}

/**
* 获取已经注册的服务
 */
func (t *RouterAll) GetRegisterServices() []*ServiceConfig {
	serviceList := make([]*ServiceConfig, 0)
	for name, vMap := range t.serviceDict { // 遍历项目
		for index := range vMap { //遍历服务

			service := t.serviceDict[name][index]
			serviceList = append(serviceList, service)
		}

	}
	return serviceList
}

/**
* 添加项目
 */
func (t *RouterAll) AddProject(project string) {
	if t.serviceDict == nil {
		t.serviceDict = make(map[string]map[string]*ServiceConfig)
	}
	serviceDict := make(map[string]*ServiceConfig)
	t.serviceDict[project] = serviceDict
}

/*
* 添加项目服务
 */
func (t *RouterAll) AddProjectService(project string, service *ServiceConfig) {
	t.serviceDict[project][service.Key] = service
}

/**
* 根据服务名获取，服务的配置
 */

func tt(tpl2 *text_template.Template) {
	params := make(map[string]interface{})
	params["nick"] = "张治"
	params["a"] = "张治"
	var buf bytes.Buffer
	tpl2.Execute(&buf, params)
	fmt.Println(buf.String())
}
func (t *RouterAll) GetProjectService(service string) (ServiceConfig, bool, string) {
	arr := strings.Split(service, ".")
	data := ServiceConfig{}
	if len(arr) != 2 {
		return data, false, "【" + service + "】服务名称格式有误"
	}
	collect_project := arr[0]
	collect_item := arr[1]
	project := t.serviceDict[collect_project]
	if project == nil {
		return data, false, "【" + collect_project + "】项目不存在"
	}

	if service_config, ok := project[collect_item]; !ok {
		return data, false, "项目：" + collect_project + "【" + collect_item + "】服务不存在"
	} else {
		//todo 可不可以先预先定义几个obj,不用临时copy

		/*这里拷贝的对象存在一个问题，如果是指针，它会重新初始化对象，但是模板挂载的方法没有了
		* 目前将指针对象，直接返回
		 */
		data = utils.Copy(*service_config).(ServiceConfig)
		return data, true, ""
	}

}

/**
* 具体项目
 */
type RouterProject struct {
	Key        string // 路由前缀
	CurrentDir string // 当前路径目录
	Path       string // 当前路径

}

/*
* 插件
 */
type Plugin struct {
	Key       string                  // 关键字
	Type      string                  // 类型
	Name      string                  // 名称
	Path      string                  // 路径
	Method    string                  // 方法
	Reg       string                  // 正则表达式
	Fields    []SubField              // 字段信息
	Enable    string                  // 是否启用
	EnableTpl *text_template.Template //启用模板
}

/*
* 项目关联服务
 */
type ProjectLink struct {
	Service    []ProjectLinkInfo // 项目路由
	CurrentDir string            // 当前路径目录
	Path       string            // 当前路径
}

/*
** 文件目录
 */
type ProjectLinkInfo struct {
	CurrentDir string // 当前路径目录
	Path       string // 当前路径
}
type ServiceList struct {
	Service    []ServiceConfig // 服务列表
	CurrentDir string          // 当前路径目录
	Path       string          // 当前路径
}

/**
* 服务
 */
type ServiceConfig struct {
	Project          string // 项目
	Name             string
	Key              string                  // 服务
	Service          string                  // 服务全称
	Params           map[string]ParamConfig  // 参数配置定义
	Module           string                  // 模块
	Http             bool                    //  是否支持http 访问
	DataFile         string                  `yaml:"data_file"` // 文件路径
	FileData         string                  // 文件内容
	Pagination       string                  //分页字段
	FileDataTpl      *text_template.Template // 文件内容的模板
	CountFile        string                  `yaml:"count_file"` // count文件路径
	CountFileData    string                  // count文件内容
	CountFileDataTpl *text_template.Template // count文件内容模板
	CurrentDir       string                  // 当前路径目录
	Path             string                  // 当前路径
	Log              bool                    // 是否写日志
	HandlerParams    []HandlerParam          `yaml:"handler_params"` //运行模块前处理参数
	ResultHandler    []HandlerParam          `yaml:"result_handler"` //运行模块完成结果处理参数
}
type ParamConfig struct {
	Name        string                  // 名称
	Template    string                  // 渲染模板
	Exec        bool                    //是否执行
	TemplateTpl *text_template.Template // 模板内容
	Type        string                  // 转换类型
	Check       SubField                //检查
	Default     interface{}             //默认值
}
type SubField struct {
	From        string                  // 来源
	To          string                  // 目标
	Name        string                  // 名称
	Template    string                  // 渲染模板
	TemplateTpl *text_template.Template // 检查内容模板
	ErrMsg      string                  `yaml:"err_msg"` // 错误提示
	ErrMsgTpl   *text_template.Template // 错误提示模板
	Field       string                  // 字段
	FieldTpl    *text_template.Template // 字段模板
	Fields      []ThirdField            // 二级字段列表
	ThirdField  string                  `yaml:"third_field"`
	ThirdFields []ThirdField            `yaml:"third_fields"` // 三级字段列表

	ThirdArrayField  string       `yaml:"third_array_field"`
	ThirdArrayFields []ThirdField `yaml:"third_array_fields"` // 三级字段列表

	Rule string // 规则
}
type ThirdField struct {
	From  string // 来源
	To    string // 目标
	Name  string // 名称
	Field string
}
type HandlerParam struct {
	Key           string                  // 服务
	Enable        string                  // 是否启用
	EnableTpl     *text_template.Template // 是否启用模板
	Name          string                  //名称
	AppendParam   bool                    `yaml:"append_param"` // 是否添加参数
	Foreach       string                  //循环数组
	Item          string                  // for 循环里面的item
	Fields        []SubField              // 子字段
	Service       map[string]interface{}  // 服务
	Template      string                  // 检查模板
	TemplateTpl   *text_template.Template //
	ErrMsg        string                  `yaml:"err_msg"` /// 错误校验
	ErrMsgTpl     *text_template.Template //
	FromList      string                  `yaml:"from_list"`   // 来源数组
	FromItem      string                  `yaml:"from_item"`   // 来源数组子项
	IfTemplate    string                  `yaml:"if_template"` // 判断2个数组是否相等
	IfTemplateTpl *text_template.Template //
	Value         string                  // 取值
	ValueTpl      *text_template.Template //
	File          string                  // 文件取值字段
	DataJson      string                  `yaml:"data_json"`   // 来源json 数据地址
	ResultName    string                  `yaml:"result_name"` // 结果参数
	SaveField     string                  `yaml:"save_field"`  // 保存字段
}
