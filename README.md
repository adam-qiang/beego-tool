<p align="center">
<a href="https://pkg.go.dev/github.com/adam-qiang/beego-tool"><img src="https://pkg.go.dev/badge/github.com/adam-qiang/beego-tool.svg" alt="Go Reference"></a>
<a href="https://en.wikipedia.org/wiki/MIT_License" rel="nofollow"><img alt="MIT" src="https://img.shields.io/badge/license-MIT-blue.svg" style="max-width:100%;"></a>
</p>

---

# beego-tool

beego框架适用工具

## 安装

``` go
go get -u github.com/adam-qiang/beego-tool
```


## 一、context

适用于beego框架的上下文工具

### 1、NewContext

创建一个新的上下文

### 2、PostForm

接收POST表单参数

### 3、Query

接收GET请求的查询参数

### 4、JsonParams

接收application/json请求头的请求参数

### 5、SetStatus

设置网络状态

### 6、SetHeader

设置响应状态

### 7、OtuPut

输出响应

### 8、OtuPutString

输出普通字符串响应

### 9、OtuPutJson

输出application/json响应

### 10、OtuPutHtml

输出HTML响应

## 二、data_tool

适用于beego框架的数据工具

### 1、ExportCsv

导出CSV

### 2、ExportExcel

数据导出excel

## 三、validate

适用于beego框架的参数校验工具

### 1、InitValidate

初始化校验（在main中进行初始化）

### 2、Valid

公共的表单校验方法