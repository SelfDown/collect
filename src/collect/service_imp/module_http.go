package collect

import (
	"bytes"
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpService struct {
	BaseHandler
}

//var db0 *sql.DB

func (s *HttpService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	configData := template.HttpConfigData

	//configData.Data
	//http.Post()

	//request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	handler := GetHandler(configData, params)
	defer handler.Close()
	// 创建请求
	result := handler.CreateRequest()
	if template.Log {
		template.LogData(handler.GetLogData())
	}
	if !result.Success {
		return result
	}
	// 自定义处理请求
	result = handler.HandlerRequest()
	if !result.Success {
		return result
	}
	result = handler.GetResult()
	if !result.Success {
		return result
	}
	if template.SuccessTpl != nil {
		r, ok := result.GetData().(map[string]interface{})
		if ok {
			success := utils.RenderTplBool(template.SuccessTpl, r)
			if !success {
				return common.NotOk(utils.Strval(r))
			}
		}
	}
	return result
}

type RequestHandler interface {
	SetConfig(*config.HttpConfig)
	GetConfig() *config.HttpConfig
	GetParams() map[string]interface{}
	SetParams(map[string]interface{})
	GetMethod() string
	GetUrl() string
	SetUrl(url string)
	GetBody() io.Reader
	SetBody(io.Reader)
	CreateRequest() *common.Result
	HandlerRequest() *common.Result
	GetResult() *common.Result
	GetLogData() map[string]interface{}
	GetData() interface{}
	SetData(data interface{})
	Close()
	//config *config.HttpConfig
}
type BaseRequestHandler struct {
	RequestHandler
	config  *config.HttpConfig
	params  map[string]interface{}
	url     string
	body    io.Reader
	data    interface{}
	timeout int64
	req     *http.Request
	resp    *http.Response
}

func (t *BaseRequestHandler) SetConfig(config *config.HttpConfig) {
	t.config = config
}
func (t *BaseRequestHandler) GetConfig() *config.HttpConfig {
	return t.config
}
func (t *BaseRequestHandler) SetParams(params map[string]interface{}) {
	t.params = params
}
func (t *BaseRequestHandler) GetParams() map[string]interface{} {
	return t.params
}
func (t *BaseRequestHandler) GetMethod() string {
	return strings.ToUpper(t.config.Method)
}

func (t *BaseRequestHandler) SetUrl(url string) {
	t.url = url
}
func (t *BaseRequestHandler) GetBody() io.Reader {
	data := t.data
	if dataStr, ok := data.(string); ok {
		return strings.NewReader(dataStr)
	}

	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	encoder.Encode(data)
	return buf
}

func (t *BaseRequestHandler) SetBody(body io.Reader) {
	t.body = body
}
func (t *BaseRequestHandler) GetData() interface{} {

	//如果是map 类型先装json字符串，进行渲染，然后给json转回来
	data := t.GetDataStr()
	p := make(map[string]interface{})
	json.Unmarshal([]byte(data), &p)
	if utils.IsValueEmpty(p) && !utils.IsValueEmpty(data) { // 处理字符串
		return p
	}
	params := t.GetParams()
	for key, value := range p {
		tpl, ok := value.(string)
		if !ok {
			continue
		}
		if !utils.IsRenderVar(tpl) {
			continue
		}
		val := utils.RenderVarOrValue(value, params)
		if val != value {
			p[key] = val
		}
	}

	return p
}
func (t *BaseRequestHandler) GetDataStr() string {
	if t.config.DataTpl == nil {
		return "{}"
	}
	data := utils.RenderTpl(t.config.DataTpl, t.params)
	return data
}
func (t *BaseRequestHandler) SetData(data interface{}) {
	t.data = data
}

func (t *BaseRequestHandler) GetLogData() map[string]interface{} {

	p := make(map[string]interface{})
	p["url"] = t.url
	p["method"] = t.config.Method
	p["data"] = t.data
	return p
}
func (t *BaseRequestHandler) Close() {
	if t.resp != nil {
		t.resp.Body.Close()
	}
}

func GetHandler(config *config.HttpConfig, params map[string]interface{}) RequestHandler {
	method := config.Method
	method = strings.ToLower(method)
	var handler RequestHandler
	if method == "get" {
		handler = &GetRequestHandler{}
	} else if method == "post" {
		handler = &PostRequestHandler{}
	} else {
		handler = &BaseRequestHandler{}
	}
	handler.SetConfig(config)
	handler.SetParams(params)
	// 设置url
	handler.SetUrl(handler.GetUrl())
	// 处理data
	handler.SetData(handler.GetData())
	// 设置body
	handler.SetBody(handler.GetBody())
	// 设置超时

	return handler
}

type GetRequestHandler struct {
	BaseRequestHandler
}

func (t *BaseRequestHandler) GetUrl() string {
	url := utils.RenderTpl(t.config.UrlTpl, t.params)
	return url
}
func (t *BaseRequestHandler) GetResult() *common.Result {
	req := t.req
	timeout := 30
	if t.config.Timeout != 0 {
		timeout = t.config.Timeout
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Do(req)

	if err != nil {
		return common.NotOk(err.Error())
	}
	t.resp = resp
	body, err := ioutil.ReadAll(resp.Body)
	resultJson := t.config.ResultJson
	// 处理不转json
	//或者状态码大于200
	if (resultJson != nil && *resultJson == false) || resp.StatusCode > 200 {
		return common.NotOk(utils.Strval(body))
	}
	result := make(map[string]interface{})
	json.Unmarshal(body, &result)
	return common.Ok(result, "请求发送成功")
}

// CreateRequest 处理请求
func (t *BaseRequestHandler) CreateRequest() *common.Result {
	var req *http.Request
	var err error
	req, err = http.NewRequest(t.GetMethod(), t.url, t.body)
	// 设置header
	if !utils.IsValueEmpty(t.config.Header) {
		for k, tpl := range t.config.HeaderTpl {
			req.Header.Set(k, utils.RenderTpl(tpl, t.params))
		}
	}
	if err != nil {
		return common.NotOk(err.Error())
	}
	//  处理拼接所有参数
	if t.config.AppendParam && !utils.IsValueEmpty(t.body) {

		jsonData, err := ioutil.ReadAll(t.body)
		if err != nil {
			return common.NotOk(err.Error())
		}
		tmp := make(map[string]interface{})
		err = json.Unmarshal(jsonData, &tmp)
		if err != nil {
			return common.NotOk(err.Error())
		}
		// 循环params 添加变量
		for k, v := range t.params {
			if _, ok := tmp[k]; !ok {
				tmp[k] = v
			}
		}
		// 重新转json
		jsonData, err = json.Marshal(tmp)
		if err != nil {
			return common.NotOk(err.Error())
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
		req.ContentLength = int64(len(jsonData))
	}
	// 设置basic auth 登录
	usernameTpl := t.config.BasicAuth.UsernameTpl
	passwordTpl := t.config.BasicAuth.PasswordTpl
	if usernameTpl != nil && passwordTpl != nil {
		username := utils.RenderTpl(usernameTpl, t.params)
		password := utils.RenderTpl(passwordTpl, t.params)
		req.SetBasicAuth(username, password)
	}

	// 设置请求
	t.req = req
	return common.Ok(req, "参数构造成功")
} // DoRequest 处理请求

// HandlerRequest 处理请求
func (t *BaseRequestHandler) HandlerRequest() *common.Result {
	return common.Ok(nil, "参数处理成功")
}

func (t *GetRequestHandler) HandlerRequest() *common.Result {
	req := t.req
	// 处理url get请求
	query := req.URL.Query()
	data := t.data
	if t.config.AppendParam {
		for k, v := range t.params {
			query.Set(k, utils.Strval(v))
		}
	}
	for k, v := range data.(map[string]interface{}) {
		query.Set(k, utils.Strval(v))
	}
	req.URL.RawQuery = query.Encode()
	return common.Ok(req, "参数构造成功")
}

func (t *GetRequestHandler) GetBody() io.Reader {

	return nil
}

type PostRequestHandler struct {
	BaseRequestHandler
}
