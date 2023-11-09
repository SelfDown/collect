package collect

import (
	"encoding/json"
)

const SuccessValue = "0"
const UnSuccessValue = "-1"

type Result struct {
	Status  int64  `json:"status"`  // 0成功，其他失败
	Count   int64  `json:"count"`   // 总数
	Success bool   `json:"success"` // 是否成功
	Code    string `json:"code"`    //编码
	Msg     string `json:"msg"`     // 消息
	Data    any    `json:"data"`    //请求参数
	finish  bool   //是否接收
}

func (r *Result) SetFinish(value bool) {
	r.finish = value
}

func (r *Result) IsFinish() bool {
	return !r.Success || r.finish
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

func (r *Result) GetCount() int64 {
	return r.Count
}

func newResult(success bool, code string, msg string, data any, count int64) *Result {
	r := Result{
		Success: success,
		Code:    code,
		Data:    data,
		Msg:     msg,
		Count:   count,
	}
	return &r
}
func OkWithCount(data any, msg string, count int64) *Result {
	r := Result{}
	return r.OkWithCount(data, msg, count)
}
func (r *Result) OkWithCount(data any, msg string, count int64) *Result {
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
