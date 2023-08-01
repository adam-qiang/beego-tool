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

## 四、数据库

适用于beego框架的数据库工具

### 1、初始化数据库（在main中进行初始化）

```golang
import _ "github.com/adam-qiang/beego-tool/database"

```

### 2、配置

```editorconfig
[mysql]
mysql_urls =
mysql_port =
mysql_user =
mysql_pass =
mysql_db =

[redis]
address =
port =
password =
database =
key =
cache_database =
cache_key =
```

- 注：redis配置中key和cache_key分别为redis普通操作前缀key和为redis缓存前缀key（可以不配置）

### 3、数据库操作

#### 1、MySQL

操作遵循beego官方操作具体见beego官方文档

#### 2、Redis

使用github.com/redis/go-redis/v9作为redis操作库进行二次封装，同时只封装了经常用到的方法，如有其他需求可自行封装可随时issue

支持以下操：

##### 2.1、KEY

- Del
- Dump
- Restore
- Exists
- Expire
- ExpireAt
- Keys
- Move
- Persist
- PExpire
- PExpireAt
- TTL
- PTTL
- Rename
- RenameNX
- Type

##### 2.2、STRING

- Set
- SetNX
- Get
- Incr
- IncrBy
- Decr
- DecrBy
- MSet
- MSetNX
- MGetMap
- StrLen

##### 2.3、HASH

- HSet
- HGet
- HGetAll
- HMSet
- HMGetMap
- HExists
- HDel
- HIncrBy
- HKeys
- HLen
- LPop
- LPush
- BLPop
- LPushX
- RPop
- RPush
- RPushX
- BRPop
- RPopLPush
- BRPopLPush
- LIndex
- LInsert
- LLen
- LRange
- LRem
- LSet

##### 2.4、SET

- SAdd
- SCard
- SDiff
- SDiffStore
- SInter
- SInterStore
- SIsMember
- SMembers
- SMove
- SRem
- SUnion
- SUnionStore

##### 2.5、SORTED SET

- ZAdd
- ZCard
- ZCount
- ZIncrBy
- ZRange
- ZRangeByScore
- ZRank
- ZRem
- ZRemRangeByRank
- ZRemRangeByScore
- ZRevRange
- ZRevRangeByLex
- ZRevRangeByScore
- ZRevRangeByScoreWithScores
- ZRevRangeWithScores
- ZRevRank
- ZScore

##### 3、Redis Cache

操作遵循beego官方操作具体见beego官方文档