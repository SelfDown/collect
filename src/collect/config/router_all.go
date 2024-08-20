package collect

import (
	"github.com/xuri/excelize/v2"
	"strings"
	text_template "text/template"

	utils "github.com/SelfDown/collect/src/collect/utils"
)

// type ModuleResult interface {
// 	Result(template *template.Template) (*common.Result, error)
// }
/**
* 项目总路由
 */
type RouterAll struct {
	Services    []RouterProject                      // 项目路由
	CurrentDir  string                               // 当前路径目录
	Path        string                               // 当前路径
	serviceDict map[string]map[string]*ServiceConfig // 服务的二维字典
	//moduleRegisterDict map[string]module_result.ModuleResult //服务注册对象
	//registerList       []module_result.ModuleResult          // 设置注册列表
	LoadStartupPlugin []Plugin          `yaml:"load_startup_plugin"` // 加载启动插件
	FileContentPlugin []Plugin          `yaml:"file_content_plugin"` // 文件内容处理插件
	ModuleHandler     []Plugin          `yaml:"module_handler"`      // 文件内容处理插件
	DataHandler       []Plugin          `yaml:"data_handler"`        // 文件内容处理插件
	BeforePlugin      []Plugin          `yaml:"before_plugin"`       // 文件内容处理插件
	AfterPlugin       []Plugin          `yaml:"after_plugin"`        // 文件内容处理插件
	moduleDict        map[string]Plugin //服务注册对象字典
}

//
//func (t *RouterAll) SetRegisterList(registerList []module_result.ModuleResult) {
//	// 设置转换字典
//	moduleDict := make(map[string]Plugin)
//	for _, module := range t.ModuleHandler {
//		moduleDict[module.Path] = module
//	}
//	for _, module := range t.DataHandler {
//		moduleDict[module.Path] = module
//	}
//	t.moduleDict = moduleDict
//	// 清空字典
//	t.moduleRegisterDict = make(map[string]module_result.ModuleResult)
//	for _, reg := range registerList {
//		// 这里根据字符串，注册层服务
//		name := reflect.TypeOf(reg).Elem().Name()
//		//moduleRegisterDict 进行赋值
//		if module, ok := moduleDict[name]; ok {
//			t.moduleRegisterDict[module.Key] = reg
//		} else {
//			log.Println("模块【" + name + "】没有注册，请检查配置！！！")
//		}
//
//	}
//
//	t.registerList = registerList
//}
//func (t *RouterAll) GetRegisterList() []module_result.ModuleResult {
//	return t.registerList
//}
//
//func (t *RouterAll) GetModuleRegister(name string) module_result.ModuleResult {
//	return t.moduleRegisterDict[name]
//}
//func (t *RouterAll) SetModuleRegister(name string, module module_result.ModuleResult) {
//	if t.moduleRegisterDict == nil {
//		t.moduleRegisterDict = make(map[string]module_result.ModuleResult, 0)
//	}
//	t.moduleRegisterDict[name] = module
//}

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
	IsHttp    bool                    //是否http
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
	Project             string // 项目
	Name                string
	Key                 string                  // 服务
	Service             string                  // 服务全称
	Params              map[string]ParamConfig  // 参数配置定义
	Schedule            Schedule                // 定时任务
	Cache               HandlerParam            // 缓存设置
	PreventDuplication  HandlerParam            `yaml:"prevent_duplication"` // 重复请求处理
	Module              string                  // 模块
	Table               string                  // 表名
	Http                bool                    //  是否支持http 访问
	Options             string                  // 可以选择字段
	MustLogin           *bool                   `yaml:"must_login"`  // 是否必须登录
	DataFile            string                  `yaml:"data_file"`   // 文件路径
	DataSource          string                  `yaml:"data_source"` // 数据源
	RunStartup          bool                    `yaml:"run_startup"` // 启动是否运行
	FileData            string                  // 文件内容
	Pagination          string                  //分页字段
	FileDataTpl         *text_template.Template // 文件内容的模板
	Count               string
	CountFile           string                  `yaml:"count_file"` // count文件路径
	CountFileData       string                  // count文件内容
	CountFileDataTpl    *text_template.Template // count文件内容模板
	CurrentDir          string                  // 当前路径目录
	Path                string                  // 当前路径
	IgnoreFields        []string                `yaml:"ignore_fields"` //忽略字段
	UpdateFields        []string                `yaml:"update_fields"` // 更新路径
	Log                 bool                    // 是否写日志
	HandlerParams       []HandlerParam          `yaml:"handler_params"` //运行模块前处理参数
	ResultHandler       []HandlerParam          `yaml:"result_handler"` //运行模块完成结果处理参数
	Filter              map[string]interface{}  //过滤条件
	ModelField          string                  `yaml:"model_field"`
	ExcelConfig         string                  `yaml:"excel_config"` // 保存路径
	ExcelConfigContent  string
	ExcelConfigData     *ExcelConfig //
	ModifyConfig        string       `yaml:"modify_config"` // 保存路径
	ModifyConfigContent string
	ModifyConfigData    *ModifyConfig //
	HttpJson            string        `yaml:"http_json"` // http路径
	HttpJsonContent     string
	DataJson            string `yaml:"data_json"` // data_json
	DataJsonContent     string
	DataJsonConfig      *DataJsonConfig // data_json 配置
	HttpConfigData      *HttpConfig     //
	Success             string
	SuccessTpl          *text_template.Template // http 请求验证模板
	Batch               HandlerParam            //批量处理参数
}
type Schedule struct {
	Enable       string                  // 是否启用
	EnableTpl    *text_template.Template // 是否启用模板
	ScheduleSpec string                  `yaml:"schedule_spec"json:"schedule_spec"`
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
	Key         string                  // 关键字
	From        string                  // 来源
	FromTpl     *text_template.Template // 来源模板
	To          string                  // 目标
	Name        string                  // 名称
	Template    string                  // 渲染模板
	Type        string                  //类型
	TemplateTpl *text_template.Template // 检查内容模板
	ErrMsg      string                  `yaml:"err_msg"json:"err_msg"` // 错误提示
	ErrMsgTpl   *text_template.Template // 错误提示模板
	Field       string                  // 字段
	FieldTpl    *text_template.Template // 字段模板
	Fields      []ThirdField            // 二级字段列表
	ThirdField  string                  `yaml:"third_field"`
	ThirdFields []ThirdField            `yaml:"third_fields"` // 三级字段列表

	ThirdArrayField  string       `yaml:"third_array_field"`
	ThirdArrayFields []ThirdField `yaml:"third_array_fields"` // 三级字段列表

	Rule      string // 规则
	Width     float64
	SaveField string `yaml:"save_field"json:"save_field"`
}
type ThirdField struct {
	From  string // 来源
	To    string // 目标
	Name  string // 名称
	Field string
}
type HandlerParam struct {
	Key                 string                  // 服务
	Enable              string                  // 是否启用
	EnableTpl           *text_template.Template // 是否启用模板
	Name                string                  //名称
	Field               string                  // 字段
	LeftField           string                  `yaml:"left_field"json:"left_field"`               // 字段
	RightField          string                  `yaml:"right_field"json:"right_field"`             // 字段
	LeftValueField      string                  `yaml:"left_value_field"json:"left_value_field"`   // 字段
	RightValueField     string                  `yaml:"right_value_field"json:"right_value_field"` // 字段
	AppendParam         bool                    `yaml:"append_param"json:"append_param"`           // 是否添加参数
	NoneFillRight       bool                    `yaml:"none_fill_right"json:"none_fill_right"`     // 是否添加右边参数
	AppendItemParam     bool                    `yaml:"append_item_param"json:"append_item_param"` // 是否添加参数
	SaveOriginal        bool                    `yaml:"save_original"json:"save_original"`         // 是否添加参数
	Foreach             string                  //循环数组
	Item                string                  // for 循环里面的item
	Fields              []SubField              // 子字段
	Service             map[string]interface{}  // 服务
	Template            string                  // 检查模板
	TemplateTpl         *text_template.Template //
	ErrMsg              string                  `yaml:"err_msg"json:"err_msg"` /// 错误校验
	ErrMsgTpl           *text_template.Template //
	FromList            string                  `yaml:"from_list"json:"from_list"`     // 来源数组
	FromItem            string                  `yaml:"from_item"json:"from_item"`     // 来源数组子项
	IfTemplate          string                  `yaml:"if_template"json:"if_template"` // 判断2个数组是否相等
	IfTemplateTpl       *text_template.Template //
	Value               string                  // 取值
	ValueTpl            *text_template.Template //
	File                string                  // 文件取值字段
	DataJson            string                  `yaml:"data_json"json:"data_json"`     // 来源json 数据地址
	ResultName          string                  `yaml:"result_name"json:"result_name"` // 结果参数
	SaveField           string                  `yaml:"save_field"json:"save_field"`   // 保存字段
	Path                string                  // 保存路径
	Params              string
	NodeNext            string                  `yaml:"node_next"json:"node_next"` // 流程流转用户
	NodeNextTpl         *text_template.Template //
	NodeKey             string                  `yaml:"node_key"json:"node_key"` // 保存字段
	IgnoreError         bool                    `yaml:"ignore_error"json:"ignore_error"`
	Type                string
	NodeType            string                  `yaml:"node_type"json:"node_type"` // 保存字段
	NodeFail            string                  `yaml:"node_fail"json:"node_fail"` // 失败流转
	NodeFailTpl         *text_template.Template //
	Rule                string
	Right               string
	Left                string
	Operation           string
	TargetTransferKey   string   `yaml:"target_transfer_key"json:"target_transfer_key"`     // 失败流转
	TargetTransferValue string   `yaml:"target_transfer_value"json:"target_transfer_value"` // 失败流转
	ValueListField      string   `yaml:"value_list_field"json:"value_list_field"`           // 失败流转
	AppendRightFields   []string `yaml:"append_right_fields"json:"append_right_fields"`     // 失败流转
	AppendLeftFields    []string `yaml:"append_left_fields"json:"append_left_fields"`       // 失败流转
	WithAddRemove       bool     `yaml:"with_add_remove"json:"with_add_remove"`             // 比较修改记录的
	Method              string
	Room                string
	Second              int64
	Children            string //children
	Pid                 string //parent_id
	Id                  string //id
	Prefix              string //flow的运行结果前缀
}
type ExcelConfig struct {
	Name   string   // 名称
	Sheets []Sheets // sheet 页
}
type ModifyConfig struct {
	Desc            string            // 备注
	OpFieldTransfer map[string]string `yaml:"op_field_transfer"json:"op_field_transfer"` // 字段冲突转换一下
	Fields          []HandlerParam
}
type DataJsonConfig struct {
	Connect  map[string]interface{}
	Finish   HandlerParam   //结束
	Services []HandlerParam // 服务
}
type HttpConfig struct {
	Url         string // 名称
	UrlTpl      *text_template.Template
	Method      string
	AppendParam bool  `json:"append_param"`
	ResultJson  *bool `json:"result_json"`
	Header      map[string]string
	HeaderTpl   map[string]*text_template.Template
	BasicAuth   BasicAuth `json:"basic_auth"`
	Data        interface{}
	Timeout     int
	DataTpl     *text_template.Template
}
type BasicAuth struct {
	Username    string
	Password    string
	UsernameTpl *text_template.Template
	PasswordTpl *text_template.Template
}
type Sheets struct {
	TitleHeight  float64 `json:"title_height"`
	Data         string
	Title        string
	TitleTpl     *text_template.Template // 是否启用模板
	Fields       []SubField              // 字段信息
	TitleStyle   excelize.Style          `json:"title_style"`   // 标题样式
	NameStyle    excelize.Style          `json:"name_style"`    // 标题样式
	ContentStyle excelize.Style          `json:"content_style"` // 标题样式
}
