/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-08 00:04:40
 */

package database

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Id  int
	Pre string
	Ng  string
	Ser string
	Th  string `orm:"size(500)"`
}

var mysqlDb = ""

func init() {
	//注册驱动,可以不加
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return
	}

	mysqlUser, _ := beego.AppConfig.String("mysql::mysql_user")
	mysqlPass, _ := beego.AppConfig.String("mysql::mysql_pass")
	mysqlUrls, _ := beego.AppConfig.String("mysql::mysql_urls")
	mysqlPort, _ := beego.AppConfig.String("mysql::mysql_port")
	mysqlDb, _ := beego.AppConfig.String("mysql::mysql_db")

	println("配置：", mysqlUser, mysqlPass, mysqlUrls, mysqlPort, mysqlDb)
	//注册 model
	orm.RegisterModel(new(Mysql))
	//注册默认数据库,必须注册一个别名为default的数据库，作为默认使用。("root:password@(127.0.0.1:3306)/databasename?charset=utf8")
	err = orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@("+mysqlUrls+":"+mysqlPort+")/"+mysqlDb+"?charset=utf8", orm.ConnMaxLifetime(600))
	if err != nil {
		return
	}
}

// IsExistTable 判断表是否存在
// @param tableName string
// @param databaseName string
// @return bool
func IsExistTable(tableName string, databaseName string) bool {
	if databaseName == "" {
		databaseName = mysqlDb
	}

	o := orm.NewOrm()
	var count int64
	err := o.Raw("SELECT COUNT(*) FROM information_schema.TABLES WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", databaseName, tableName).QueryRow(&count)
	if err == nil && count == 0 {
		return false
	}
	return true
}
