/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-07 00:40:03
 */

package tool

import (
	"encoding/json"
	"fmt"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"net/http"
)

type Content map[string]interface{}

// Context 上下文结构
type Context struct {
	Writer     *beegoContext.Response
	Req        *beegoContext.Context
	Path       string
	Method     string
	StatusCode int
}

// ReturnMsg
//
//	@Description: 响应结构体
type ReturnMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewContext 创建新的上下文
// @param ctx *beegoContext.Context
// @return *Context
func NewContext(ctx *beegoContext.Context) *Context {
	return &Context{
		Writer: ctx.ResponseWriter,
		Req:    ctx,
		Path:   ctx.Request.URL.Path,
		Method: ctx.Request.Method,
	}
}

// PostForm 获取post表单参数
// @receiver ctx *Context
// @param key string
// @return string
func (ctx *Context) PostForm(key string) string {
	return ctx.Req.Request.FormValue(key)
}

// Query 查询get请求参数
// @receiver ctx *Context
// @param key string
// @return string
func (ctx *Context) Query(key string) string {
	return ctx.Req.Request.URL.Query().Get(key)
}

// JsonParams 接收application/json请求头的请求参数
// @receiver ctx *Context
// @param params map[string]interface{}
func (ctx *Context) JsonParams(params map[string]interface{}) {
	err := json.Unmarshal(ctx.Req.Input.RequestBody, &params)
	if err != nil {
		return
	}
}

// SetStatus 设置网络状态
// @receiver ctx *Context
// @param code int
func (ctx *Context) SetStatus(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

// SetHeader 设置响应header
// @receiver ctx *Context
// @param key string
// @param value string
func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

// OtuPut 响应输出
// @receiver ctx *Context
// @param code int
// @param data []byte
func (ctx *Context) OtuPut(code int, data []byte) {
	ctx.SetStatus(code)
	_, err := ctx.Writer.Write(data)
	if err != nil {
		return
	}
}

// OtuPutString 输出字符串
// @receiver ctx *Context
// @param code int
// @param format string
// @param values ...interface{}
func (ctx *Context) OtuPutString(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.SetStatus(code)
	_, err := ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		return
	}
}

// OtuPutJson 输出json
// @receiver ctx *Context
// @param code int
// @param obj interface{}
func (ctx *Context) OtuPutJson(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetStatus(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

// OtuPutHtml 输出HTML响应
// @receiver ctx *Context
// @param code int
// @param html string
func (ctx *Context) OtuPutHtml(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.SetStatus(code)
	_, err := ctx.Writer.Write([]byte(html))
	if err != nil {
		return
	}
}
