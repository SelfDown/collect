package collect

import (
	"bytes"
	"fmt"
	common "github.com/SelfDown/collect/src/collect/common"
	"github.com/SelfDown/collect/src/collect/config"
	startup "github.com/SelfDown/collect/src/collect/startup"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"mime/multipart"
	"net/http"
	"reflect"
	"sync"
	text_template "text/template"
	"time"
	"unicode/utf8"
	"log"
)

type TemplateService struct {
	OpUser           string
	session          *sessions.Session // 设置session
	context          *gin.Context      // 设置http 上下文
	Request          interface{}
	IsFileResponse   bool
	ResponseFilePath string
	ResponseFileName string
	File             multipart.File        // 单个上传文件
	FileHeader       *multipart.FileHeader // 单个上传文件
	thirdData        map[string]interface{}
	ws               *websocket.Conn
	ts               TsFile
}
type TsFile interface {
	ReadAt(b []byte, off int64) (int, error)
	Size() int64
}

/*
* 数据库表
 */
var _model DatabaseModel

type DatabaseModel interface {
	GetModel(tableName string) interface{}
	CloneModel(tableName string) interface{}
	GetPrimaryKey(tableName string) []string
}

func (s *TemplateService) SetTsFile(ts TsFile) {
	s.ts = ts
}
func (s *TemplateService) GetTsFile() TsFile {
	return s.ts
}
func (s *TemplateService) SetWs(ws *websocket.Conn) {
	s.ws = ws
}
func (s *TemplateService) GetWs() *websocket.Conn {
	return s.ws
}

func (s *TemplateService) GetWsOutput(params map[string]interface{}) WSoutput {

	return WSoutput{
		ws: s.ws,
		params: params,
		ts: s,
	}
}

// 扩展模块设置
func (s *TemplateService) SetThirdData(key string, value interface{}) {
	if s.thirdData == nil {
		s.thirdData = make(map[string]interface{})
	}
	s.thirdData[key] = value
}
func (s *TemplateService) HasThirdData(key string) bool {
	_, has := s.thirdData[key]
	return has
}

//扩展模块获取
func (s *TemplateService) GetThirdData(key string) interface{} {
	return s.thirdData[key]
}

//扩展模块获取
func (s *TemplateService) RemoveThirdData(key string) {
	delete(s.thirdData, key)
}

// SetDatabaseModel 设置数据库表
func SetDatabaseModel(m DatabaseModel) {
	_model = m
}
func (*TemplateService) GetModel(tableName string) interface{} {
	return _model.GetModel(tableName)
}
func (*TemplateService) CloneModel(tableName string) interface{} {
	return _model.CloneModel(tableName)
}
func (*TemplateService) GetPrimaryKey(tableName string) []string {
	return _model.GetPrimaryKey(tableName)
}

// RunScheduleService 添加定时任务
func RunScheduleService() []*collect.ServiceConfig {
	services := collect.GetLocalRouter().GetRegisterServices()
	scheduleService := make([]*collect.ServiceConfig, 0)
	params := make(map[string]interface{})
	c := cron.New()
	i := 0
	for _, service := range services {
		if !utils.IsValueEmpty(service.Schedule.Enable) && !utils.IsValueEmpty(service.Schedule.ScheduleSpec) {
			run := utils.RenderTplBool(service.Schedule.EnableTpl, params)
			if run {
				scheduleService = append(scheduleService, service)
				i++
				paramService := make(map[string]interface{})
				paramService["service"] = service.Service
				fmt.Println(utils.Strval(i) + ".添加定时任务[" + service.Service + "]" + service.Schedule.ScheduleSpec)
				c.AddFunc(service.Schedule.ScheduleSpec, func() {
					ts := TemplateService{OpUser: "schedule"}
					ts.ResultInner(paramService)
				})
			}
		}

	}
	if len(scheduleService) > 0 {
		c.Start()
	}
	return scheduleService
}

const batchSize = 100 // 每批次处理的消息数量
var queue chan Msg

type Msg struct {
	DataType string
	Data     interface{}
	Params     map[string]interface{}
	CreateTime int64
}

// RunStartupService 添加启动服务
func RunStartupService() []*collect.ServiceConfig {
	services := collect.GetLocalRouter().GetRegisterServices()
	startupService := make([]*collect.ServiceConfig, 0)
	i := 0
	for _, service := range services {
		if service.RunStartup {
			startupService = append(startupService, service)
			i++
			paramService := make(map[string]interface{})
			paramService["service"] = service.Service
			fmt.Println(utils.Strval(i) + ".执行启动服务[" + service.Service + "]" + service.Schedule.ScheduleSpec)
			ts := TemplateService{OpUser: "startup"}
			ts.ResultInner(paramService)
		}

	}

	var wg sync.WaitGroup
	queue = make(chan Msg, batchSize)
	wg.Add(1)
	go consumer(queue, &wg)
	return startupService
}

func updateTimer(timer *time.Timer) {
	timer.Reset(time.Second * 10)
}
// consumer 函数从队列中接收并处理指定数量的消息
func consumer(queue <-chan Msg, wg *sync.WaitGroup) {
	defer wg.Done()
	batch := make([]Msg, 0, batchSize)
	timer := time.NewTimer(time.Second * 10)
	updateTime := utils.Debounce(func() {
		updateTimer(timer)
	}, 5)
	for {
		select {
		case msg, ok := <-queue:
			updateTime()
			if !ok {
				// 通道已关闭，退出循环
				return
			}
			batch = append(batch, msg)
			if len(batch) == batchSize {
				// 处理当前批次的消息
				processBatch(batch)
				batch = batch[:0] // 清空批次切片以便接收下一批消息
			}
		case <-timer.C:
			if len(batch)>0{
				processBatch(batch)
				batch = batch[:0] // 清空批次切片以便接收下一批消息	
			}
			
		default:
			// 如果没有消息可接收，则休眠以减少CPU使用率
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// processBatch 函数处理一个批次的消息
func processBatch(batch []Msg) {
	paramService:=make(map[string]interface{})
	dataList:=make([]map[string]interface{},0)
	for _, msg := range batch {
		item:=make(map[string]interface{},0)
		item["data_type"]=msg.DataType
		item["data"]=msg.Data
		item["params"]=msg.Params
		item["create_time"]=msg.CreateTime
		dataList = append(dataList, item)
	}
	ts := TemplateService{OpUser: "msg"}
	paramService["service"]=utils.GetAppKey("msg_service")
	paramService["data_list"]=dataList
	r:=ts.ResultInner(paramService)
	if !r.Success{
		log.Println(r.GetMsg())
	}
	
}

var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
type WSoutput struct {
	ws *websocket.Conn
	params map[string]interface{}
	ts *TemplateService
}

// Write: implement Write interface to write bytes from ssh server into bytes.Buffer.
func (w *WSoutput) Write(p []byte) (int, error) {
	// 处理非utf8字符
	if !utf8.Valid(p) {
		bufStr := string(p)
		buf := make([]rune, 0, len(bufStr))
		for _, r := range bufStr {
			if r == utf8.RuneError {
				buf = append(buf, []rune("@")...)
			} else {
				buf = append(buf, r)
			}
		}

		p = []byte(string(buf))
	}
	err := w.ws.WriteMessage(websocket.TextMessage, p)
	w.ts.AddMsg("term", string(p),w.params)
	// 添加日志消息
	
	return len(p), err
}

func HandlerWsRequest(context *gin.Context) {
	// 如果后面还有ws的任务，可以将ws的配置，移动到这里，目前只有一个就写在了shell_term 中
	ts := getTs(context)

	w := context.Writer
	req := context.Request
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		//log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ts.SetWs(c)
	//var data *common.Result
	params := make(map[string]interface{})
	params["service"] = utils.GetAppKey("ws_service")
	params["token"] = context.Param("token")
	utils.Block{
		Try: func() {
			r := ts.Result(params, true)
			if !r.Success {
				c.WriteMessage(websocket.TextMessage, []byte(r.GetMsg()))
			}

		},
		Catch: func(e utils.Exception) {
			dv := reflect.ValueOf(e)
			_ = common.NotOk(dv.String())
		},
	}.Do()

}
func getTs(c *gin.Context) TemplateService {
	s := sessions.Default(c)

	user_id_key := utils.GetAppKey("user_id_key")
	// session 中设置用户ID
	userId := s.Get(user_id_key)
	var opUser string
	if userId != nil {
		opUser = userId.(string)
	} else {
		opUser = ""
	}
	ts := TemplateService{OpUser: opUser}
	ts.SetSession(&s)
	ts.SetContext(c)
	return ts
}
func HandlerRequest(c *gin.Context) {
	ts := getTs(c)

	//设置参数
	params := make(map[string]interface{})
	c.Bind(&params)
	if c.Request.PostForm != nil {
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				params[k] = v[0]
			} else {
				params[k] = nil
			}

		}
		// 处理单个文件上传
		file, fileHeader, error := c.Request.FormFile("file")
		if error != nil {
			data := common.NotOk(error.Error())
			c.JSON(200, data)
		}

		ts.FileHeader = fileHeader
		ts.File = file
	}
	// 设置session

	// 处理结果
	var data *common.Result
	utils.Block{
		Try: func() {
			data = ts.Result(params, true)
		},
		Catch: func(e utils.Exception) {
			dv := reflect.ValueOf(e)
			data = common.NotOk(dv.String())
		},
	}.Do()

	// 处理amis结果
	//handlerAmis(data)
	if ts.IsFileResponse {

		if !utils.IsValueEmpty(ts.ResponseFilePath) {
			filename := ts.ResponseFileName
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
			c.Writer.Header().Add("Content-Type", "application/octet-stream")
			c.File(ts.ResponseFilePath)
		}

	} else {
		c.JSON(200, data)
	}
}

func init() {
	// 加载配置
	// 加载系统插件，主要加载count、file_data,
	routerAll := startup.LoadSystemServices()
	//获取启动注册的服务列表，然后路由设置注册服务
	SetRegisterList(&routerAll, GetRegisterList())
	//设置服务
	collect.SetLocalRouter(routerAll)

}
func (t *TemplateService) SetSession(session *sessions.Session) {

	t.session = session
}

func (t *TemplateService) AddMsg(dataType string,data interface{},params map[string]interface{}) {
	now :=time.Now().UnixNano()
	msg:=Msg{DataType: dataType,Data: data,Params: params,CreateTime:now }
	queue<-msg
}


func (t *TemplateService) SetContext(context *gin.Context) {

	t.context = context
}

func (t *TemplateService) GetContext() *gin.Context {
	return t.context
}

// GetSession 获取session
func (t *TemplateService) GetSession() *sessions.Session {
	return t.session
}

func (t *TemplateService) getModuleResultObj(moduleName string) ModuleResult {
	// 根据模块名称获取，模块对象
	module := GetModuleRegister(moduleName)
	return module

}
func IsPluginEnable(Tpl *text_template.Template, plugin collect.Plugin) bool {
	var buf bytes.Buffer

	err := Tpl.Execute(&buf, plugin)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return gocast.ToBool(buf.String())

}
func (t *TemplateService) before(params map[string]interface{}, isHttp bool) (*collect.Template, *common.Result) {
	serviceName := collect.GetServiceName(params)
	if utils.IsValueEmpty(serviceName) {
		errMsg := "请求参数【service】不存在，请检查传入参数"
		return nil, common.NotOk(errMsg)
	}

	// 根据service 名称获取配置
	cfg := collect.NewTemplateService(params)
	if !cfg.Success {
		return nil, cfg
	}
	// 生成模板
	temp := collect.Template{}
	// 设置服务配置
	temp.ServiceConfig = cfg.GetData().(collect.ServiceConfig)
	// 设置全局路由配置
	temp.RouterAllConfig = collect.GetLocalRouter()

	// 设置参数
	temp.SetParams(params)
	// todo: 这里只示例了一个用户ID
	// 设置操作用户,需要将模块的变量，赋值给temp,比如session,event_id，http 的请求对象，http 的请求头
	// 内部的服务调用，也是如此，比如template 生成了一个事件ID,后面服务沿用这个事件ID,直到服务就结束
	temp.OpUser = t.OpUser
	if temp.Log {
		msg := "【" + temp.OpUser + "】访问:" + serviceName
		temp.LogData(msg)
		temp.LogData(params)
	}
	// 如果是http 请求，并且配置必须登录，如果没有用户ID,则提示必须登录
	mustLoginConfig := utils.GetAppKeyWithDefault("must_login", "true")
	mustLogin := true
	// 总开关
	if mustLoginConfig != "true" {
		mustLogin = false
	}
	// 服务开关
	if temp.MustLogin != nil && *temp.MustLogin == false {
		mustLogin = false
	}
	// http 登录判断
	if isHttp && !temp.Http {
		errMsg := serviceName + "不支持http 访问"
		return nil, common.NotOk(errMsg)
	}
	// 用户登录判断
	if isHttp && mustLogin && utils.IsValueEmpty(t.OpUser) {
		errMsg := "请登录！！！"
		return nil, common.NotOk(errMsg)
	}
	var loader BeforeLoader
	for _, plugin := range temp.GetBeforePlugins() {
		//插件是否启用
		if !IsPluginEnable(plugin.EnableTpl, plugin) {
			continue
		}
		pluginResult := collect.CallPluginFunc(&loader, plugin, &temp, t)
		if pluginResult.IsFinish() {
			return &temp, pluginResult
		}

	}

	return &temp, common.Ok(&temp, "成功")
}
func ExecTime(template *collect.Template, start time.Time, method string) {
	dis := time.Now().Sub(start).Seconds()
	template.LogData("服务：" + template.GetService() + " [" + method + "]耗时：" + utils.Strval(dis) + "s")
}

func (t *TemplateService) execute(temp *collect.Template) *common.Result {
	if temp.Log {
		defer ExecTime(temp, time.Now(), temp.Module)
	}
	// 调用模块结果
	result := t.getModuleResultObj(temp.Module)
	params := temp.GetParams()
	if result == nil {
		errMsg := "module【" + temp.Module + "】模块不存在，请检查配置"
		temp.LogErr(errMsg)
		temp.LogErr(params)
		return common.NotOk(errMsg)
	}
	data := result.Result(temp, t)

	return data
}
func (t *TemplateService) after(temp *collect.Template) *common.Result {

	var loader AfterLoader
	for _, plugin := range temp.GetAfterPlugins() {
		//插件是否启用
		if !IsPluginEnable(plugin.EnableTpl, plugin) {
			continue
		}
		pluginResult := collect.CallPluginFunc(&loader, plugin, temp, t)
		if !pluginResult.Success {
			return pluginResult
		}

	}

	return common.Ok(&temp, "成功")
}
func (t *TemplateService) ResultInner(params map[string]interface{}) *common.Result {
	var data *common.Result
	utils.Block{
		Try: func() {
			copyMap := utils.Copy(params).(map[string]interface{})
			data = t.Result(copyMap, false)
		},
		Catch: func(exception utils.Exception) {
			data = common.NotOk(utils.Strval(exception))
		},
	}.Do()
	return data
}
func (t *TemplateService) Result(params map[string]interface{}, isHttp bool) *common.Result {

	// 执行处理前

	temp, beforeResult := t.before(params, isHttp)
	//如果失败了，或则结束直接返回
	if beforeResult.IsFinish() {
		return beforeResult
	}
	// 处理中
	data := t.execute(temp)
	// 设置结果
	// 如果参数处理中没有设置过结果，则设置模块处理的结果 ，优先参数返回的结果，
	// 一般是空模块参数中配置了结果，如果实在参数中配置了结果，模块中也需要结果，那么参数中的结果不能配置
	// 最好result_handler 配置结果
	if !temp.HasResult() {
		temp.SetResult(data)
	}

	if !data.Success {
		return data
	}
	afterResult := t.after(temp)
	if !afterResult.Success {
		return afterResult
	}
	// 获取结果
	result := temp.GetResult()

	return result
}
