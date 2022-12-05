package collect

import (
	"encoding/json"
)

const SuccessValue = "0"
const UnSuccessValue = "-1"

type Result struct {
	Success bool   `json:"success"` // 是否成功
	Code    string `json:"code"`    //编码
	Msg     string `json:"msg"`     // 消息
	Data    any    `json:"data"`    //请求参数
	Count   int    `json:"count"`   // 总数
}

func (r *Result) GetSuccess() bool {

	return r.Success
}

func (r *Result) GetCode() string {
	return r.Code
}

func (r *Result) GetData() any {
	return r.Data
}

func (r *Result) GetMsg() string {
	return r.Msg
}

func (r *Result) GetCount() int {
	return r.Count
}

func newResult(success bool, code string, msg string, data any, count int) *Result {
	r := Result{
		Success: success,
		Code:    code,
		Data:    data,
		Msg:     msg,
		Count:   count,
	}
	return &r
}
func (r *Result) OkWithCount(data any, msg string, count int) *Result {
	result := newResult(true, SuccessValue, msg, data, count)
	return result
}
func Ok(data any, msg string) *Result {
	r := Result{}
	return r.Ok(data, msg)
}
func (r *Result) Ok(data any, msg string) *Result {
	result := newResult(true, SuccessValue, msg, data, 0)
	return result
}

func NotOk(msg string) *Result {
	r := Result{}
	return r.NotOk(msg)
}
func (r *Result) NotOk(msg string) *Result {
	result := newResult(false, UnSuccessValue, msg, nil, 0)
	return result
}

func (r *Result) ToString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
