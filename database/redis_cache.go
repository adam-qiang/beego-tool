/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-08-01 00:10:27
 */

package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	beego "github.com/beego/beego/v2/server/web"
)

var RedisCache cache.Cache

func init() {
	redisHost, _ := beego.AppConfig.String("redis::address")
	port, _ := beego.AppConfig.String("redis::port")
	dataBase, _ := beego.AppConfig.String("redis::cache_database")
	password, _ := beego.AppConfig.String("redis::password")
	redisKey, _ := beego.AppConfig.String("redis::cache_key")

	config := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%s","password":"%s"}`, redisKey, redisHost+":"+port, dataBase, password)
	var err error
	RedisCache, err = cache.NewCache("redis", config)
	if err != nil || RedisCache == nil {
		errMsg := "failed to init redis cache"
		panic(errMsg)
	}

}
