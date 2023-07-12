/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-07 00:44:14
 */

package tool

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"reflect"
	"strings"
	"time"
)

// MessageTemples 错误提示模板
var MessageTemples = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小为%d",
	"Max":          "最大为%d",
	"Range":        "范围在%d至%d",
	"MinSize":      "最小长度为%d",
	"MaxSize":      "最大长度为%d",
	"Length":       "长度必须是%d",
	"Alpha":        "必须是有效的字母字符",
	"Numeric":      "必须是有效的数字字符",
	"AlphaNumeric": "必须是有效的字母或数字字符",
	"Match":        "必须匹配格式%s",
	"NoMatch":      "必须不匹配格式%s",
	"AlphaDash":    "必须是有效的字母或数字或破折号(-_)字符",
	"Email":        "必须是有效的邮件地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "必须是有效的base64字符",
	"Mobile":       "必须是有效手机号码",
	"Tel":          "必须是有效电话号码",
	"Phone":        "必须是有效的电话号码或者手机号码",
	"ZipCode":      "必须是有效的邮政编码",
}

// InitValidate
// 初始化校验（在main中进行初始化）
func InitValidate() {
	_setDefaultMessage()
}

// Valid 公共的表单校验方法
// @param obj interface{}
// @return error string
func Valid(obj interface{}, validate interface{}) (error string) {
	valid := validation.Validation{}
	b, _ := valid.Valid(obj)
	if !b {
		//通过反射获取结构体
		st := reflect.TypeOf(validate)

		for _, err := range valid.Errors {
			//获取验证的字段名和提示信息的别名
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			//返回验证的错误信息
			return strings.Replace(err.Message, err.Field, alias, 1)
		}
	}

	return ""
}

// SetDefaultMessage
//
//	默认设置通用的错误验证和提示项
func _setDefaultMessage() {
	if len(MessageTemples) == 0 {
		return
	}
	//将默认的提示信息转为自定义
	for k, _ := range MessageTemples {
		validation.MessageTmpls[k] = MessageTemples[k]
	}

	//增加默认的自定义验证方法
	_ = validation.AddCustomFunc("DateFormat", _dateFormat)
	_ = validation.AddCustomFunc("DateTimeFormat", _dateTimeFormat)
	_ = validation.AddCustomFunc("Duration", _duration)
}

// _dateFormat 校验日期格式是否是Y-m-d
var _dateFormat validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
	//断言传入的是字符串
	dateObj, err := obj.(string)
	if !err {
		v.AddError(key, "必须是有效的日期格式（Y-m-d）")
		return
	}
	dateTmp := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	_, ok := time.ParseInLocation(dateTmp, dateObj, loc)
	if ok != nil {
		v.AddError(key, "必须是有效的日期格式（Y-m-d）")
		return
	}
	return
}

// _dateTimeFormat 校验时间格式是否是Y-m-d H:i:s
var _dateTimeFormat validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
	//断言传入的是字符串
	dateObj, err := obj.(string)
	if !err {
		v.AddError(key, "必须是有效的时间格式（Y-m-d H:i:s）")
		return
	}

	dateTmp := "2006-01-02 15:04:05"
	localTime, ok := time.ParseInLocation(dateTmp, dateObj, time.Local)
	fmt.Println(localTime)
	if ok != nil {
		v.AddError(key, "必须是有效的时间格式（Y-m-d H:i:s）")
		return
	}
	return
}

// _duration 工作时长校验
var _duration validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
	switch obj.(type) {
	case float32:
		objFloat := obj.(float32)
		if objFloat < 0.1 {
			v.AddError(key, "最小为0.1")
			return
		}
		if objFloat > 24.0 {
			v.AddError(key, "最大为24.0")
			return
		}
	default:
		v.AddError(key, "数据类型必须是float32")
	}
}
