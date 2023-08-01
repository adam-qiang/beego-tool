/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-08-01 00:20:05
 */

package database

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()
var redisKey = ""

func init() {
	redisHost, _ := beego.AppConfig.String("redis::address")
	port, _ := beego.AppConfig.String("redis::port")
	dataBase, _ := beego.AppConfig.String("redis::database")
	dataBaseNum, _ := strconv.Atoi(dataBase)
	password, _ := beego.AppConfig.String("redis::password")
	redisKey, _ = beego.AppConfig.String("redis::cache_key")

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + port,
		Password: password,    // no password set
		DB:       dataBaseNum, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("failed to init redis：" + err.Error())
	}
}

// Del 删除一个指定key
// @param key string
// @return bool
func Del(key string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	result := rdb.Del(ctx, key).Val()

	if result == 1 {
		return true
	}
	return false
}

// Dump 序列化给定key，并返回被序列化的值，使用Restore命令可以将这个值反序列化为Redis键
// @param key string
// @return bool
func Dump(key string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Dump(ctx, key).Val()
}

// Restore 反序列化给定的序列化值，并将它和给定的key关联
// @param key string
// @param ttl int64
// @param value string
// @return string
func Restore(key string, ttl int64, value string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Restore(ctx, key, time.Duration(ttl)*time.Second, value).Val()
}

// Exists 判断一个指定key是否存在
// @param key string
// @return bool
func Exists(key string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	result := rdb.Exists(ctx, key).Val()

	if result == 1 {
		return true
	}
	return false
}

// Expire 设置一个指定key的过期时间
// @param key string
// @param expiration int64
// @return bool
func Expire(key string, expiration int64) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.Expire(ctx, key, time.Duration(expiration)*time.Second).Err()
	if err == nil {
		return true
	}
	return false
}

// ExpireAt 与Expire类似，都用于为key设置生存时间。但ExpireAt接受的时间参数是 UNIX 时间戳(unix timestamp)
// @param key string
// @param  timestamp time.Time
// @return bool
func ExpireAt(key string, timestamp time.Time) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.ExpireAt(ctx, key, timestamp).Err()
	if err == nil {
		return true
	}
	return false
}

// Keys 查找所有符合给定模式 pattern 的 key
// * 匹配数据库中所有 key 。
// h?llo 匹配hello，hallo和hxllo等。
// h*llo 匹配 hllo和heeeeello等。
// h[ae]llo 匹配hello和hallo，但不匹配 hillo
// @param pattern string
// @param []]string
func Keys(pattern string) []string {
	if redisKey != "" {
		pattern = redisKey + ":" + pattern
	}

	return rdb.Keys(ctx, pattern).Val()
}

// Move 将当前数据库的key移动到给定的数据库db当中
// @param key string
// @param db int
// @return bool
func Move(key string, db int) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Move(ctx, key, db).Val()
}

// Persist 移除给定 key 的生存时间，将这个 key 从『易失的』(带生存时间 key )转换成『持久的』(一个不带生存时间、永不过期的 key )
// @param key string
// @return bool
func Persist(key string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Persist(ctx, key).Val()
}

// PExpire 与Expire作用类似，但是它以毫秒为单位设置key的生存时间，而不像Expire以秒为单位
// @param key string
// @param timeout time.Duration
// @return bool
func PExpire(key string, timeout time.Duration) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.PExpire(ctx, key, timeout).Val()
}

// PExpireAt  与ExpireAt命令类似，但它以毫秒为单位设置key的过期unix时间戳，而不是像ExpireAt以秒为单位
func PExpireAt(key string, timestamp time.Time) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.PExpireAt(ctx, key, timestamp).Err()
	if err == nil {
		return true
	}
	return false
}

// TTL 获取一个指定key的过期时间
// @param key string
// @return int64
func TTL(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	ttl := rdb.TTL(ctx, key).Val().Seconds()

	return int64(ttl)
}

// PTTL 与TTL类似，但它以毫秒为单位返回key剩余生存时间，而不是像TTL以秒为单位
// @param key string
// @return int64
func PTTL(key string) time.Duration {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.PTTL(ctx, key).Val()
}

// Rename 将key改名为newKey
// @param key string
// @param newKey string
// @return bool
func Rename(key, newKey string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
		newKey = redisKey + ":" + newKey
	}

	err := rdb.Rename(ctx, key, newKey).Err()
	if err == nil {
		return true
	}
	return false
}

// RenameNX 当且仅当newKey不存在时，将key改名为newKey
// @param key string
// @param newKey string
// @return bool
func RenameNX(key, newKey string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
		newKey = redisKey + ":" + newKey
	}

	return rdb.RenameNX(ctx, key, newKey).Val()
}

// Type 返回 key 所储存的值的类型
// @param key string
// @return string
func Type(key string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Type(ctx, key).Val()
}

// Set 给指定key设置value
// @param key string
// @param value string
// @param expiration int64
// @return bool
func Set(key string, value interface{}, expiration int64) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	if err == nil {
		return true
	}
	return false
}

// SetNX 给指定key设置value，当且仅当 key 不存在
// @param key string
// @param value string
// @param expiration int64
// @return bool
func SetNX(key string, value interface{}, expiration int64) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.SetNX(ctx, key, value, time.Duration(expiration)*time.Second).Val()
}

// Get 获取一个指定key
// @param key string
// @return string
func Get(key string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Get(ctx, key).Val()
}

// Incr 将key中储存的数字值增一
// @param key string
// @return int64
// @return bool
func Incr(key string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Incr(ctx, key).Result()
}

// IncrBy 将key中储存的数字值增加increment
// @param key string
// @param increment int64
// @return int64
// @return bool
func IncrBy(key string, increment int64) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.IncrBy(ctx, key, increment).Result()
}

// Decr 将key中储存的数字值减一
// @param key string
// @return int64
// @return bool
func Decr(key string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.Decr(ctx, key).Result()
}

// DecrBy 将key中储存的数字值减少increment
// @param key string
// @param increment int64
// @return int64
// @return bool
func DecrBy(key string, increment int64) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.DecrBy(ctx, key, increment).Result()
}

// MSet 同时设置一个或多个key-value对
// @param keyValues map[string]string
// @return bool
func MSet(keyValues map[string]interface{}) bool {
	newKeyValues := make(map[string]interface{})
	if redisKey != "" {
		for k, v := range keyValues {
			newKeyValues[redisKey+":"+k] = v
			delete(keyValues, k)
		}
	} else {
		newKeyValues = keyValues
	}

	err := rdb.MSet(ctx, newKeyValues).Err()
	if err == nil {
		return true
	}
	return false
}

// MSetNX 同时设置一个或多个key-value对，当且仅当所有给定 key 都不存在
// @param keyValues map[string]string
// @return bool
func MSetNX(keyValues map[string]interface{}) bool {
	newKeyValues := make(map[string]interface{})
	if redisKey != "" {
		for k, v := range keyValues {
			newKeyValues[redisKey+":"+k] = v
			delete(keyValues, k)
		}
	} else {
		newKeyValues = keyValues
	}

	return rdb.MSetNX(ctx, newKeyValues).Val()
}

// MGetMap MGet以map数据类型返回所有(一个或多个)给定key的值
// @param keys []string
// @return map[string]interface{}
func MGetMap(keys []string) map[string]interface{} {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	result := rdb.MGet(ctx, keys...).Val()

	list := make(map[string]interface{})
	if len(result) == 0 {
		return list
	}
	for k, v := range result {
		key := keys[k]
		if redisKey != "" {
			key = strings.Replace(key, redisKey+":", "", 1)
		}

		if v == nil {
			list[key] = ""
		} else {
			list[key] = v
		}
	}
	return list
}

// StrLen 返回key所储存的字符串值的长度
// @param key string
// @return int64
func StrLen(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.StrLen(ctx, key).Val()
}

// HSet 设置一个hash类型key的field的值
// @param key string
// @param field string
// @param value string
// @return bool
func HSet(key, field string, value interface{}) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.HSet(ctx, key, field, value).Err()
	if err == nil {
		return true
	}
	return false
}

// HGet 获取一个hash类型key的field的值
// @param key string
// @param field string
// @return string
func HGet(key, field string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HGet(ctx, key, field).Val()
}

// HGetAll 获取一个hash类型key的所有field和value
// @param key string
// @return map[string]string
func HGetAll(key string) map[string]string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HGetAll(ctx, key).Val()
}

// HMSet 设置一个hash类型key的多个field和value
// @param key string
// @param fieldValues map[string]string
// @return bool
func HMSet(key string, fieldValues map[string]interface{}) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.HMSet(ctx, key, fieldValues).Err()
	if err == nil {
		return true
	}
	return false
}

// HMGetMap HMGet以map数据类型获取一个hash类型key的多个field的值
// @param key string
// @param fields []string
// @return map[string]interface{}
func HMGetMap(key string, fields []string) map[string]interface{} {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	result := rdb.HMGet(ctx, key, fields...).Val()

	list := make(map[string]interface{})
	if len(result) == 0 {
		return list
	}
	for k, v := range result {
		if v == nil {
			list[fields[k]] = ""
		}
		list[fields[k]] = v
	}
	return list
}

// HExists 判断一个hash类型key的field是否存在
// @param key string
// @param field string
// @return bool
func HExists(key, field string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HExists(ctx, key, field).Val()
}

// HDel 删除一个hash类型key的field
// @param key string
// @param fields []string
// @return bool
func HDel(key string, fields []string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	result := rdb.HDel(ctx, key, fields...).Val()

	if result == 1 {
		return true
	}
	return false
}

// HIncrBy 增加一个hash类型key的field的值
// @param key string
// @param field string
// @param incr int64
// @return int64
// @return bool
func HIncrBy(key, field string, incr int64) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HIncrBy(ctx, key, field, incr).Result()
}

// HKeys 获取一个hash类型key的所有field
// @param key string
// @return []string
func HKeys(key string) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HKeys(ctx, key).Val()
}

// HLen 获取一个hash类型key的field数量
// @param key int64
func HLen(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.HLen(ctx, key).Val()
}

// LPop 从左侧移出并获取列表的第一个元素
// @param key string
// @return string
func LPop(key string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LPop(ctx, key).Val()
}

// LPush 向列表左侧添加元素
// @param key string
// @param value string
// @return int64
// @return bool
func LPush(key, value string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LPush(ctx, key, value).Result()
}

// BLPop LPop的阻塞式弹出（从左侧）
func BLPop(key string, timeout int64) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.BLPop(ctx, time.Duration(timeout)*time.Second, key).Val()
}

// LPushX 向列表左侧添加元素，仅当列表中不存在该元素时，才插入
// @param key string
// @param value string
// @return int64
// @return bool
func LPushX(key, value string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LPushX(ctx, key, value).Result()
}

// RPop 从右侧移出并获取列表的第一个元素
// @param key string
// @return string
func RPop(key string) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.RPop(ctx, key).Val()
}

// RPush 向列表右侧添加元素
// @param key string
// @param value string
// @return int64
// @return error
func RPush(key, value string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.RPush(ctx, key, value).Result()
}

// RPushX 向列表左侧添加元素，仅当列表中不存在该元素时，才插入
// @param key string
// @param value string
// @return int64
// @return bool
func RPushX(key, value string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.RPushX(ctx, key, value).Result()
}

// BRPop RPop的阻塞式弹出（从右侧）
func BRPop(key string, timeout int64) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.BRPop(ctx, time.Duration(timeout)*time.Second, key).Val()
}

// RPopLPush 在一个原子时间内，执行以下两个动作：
// 1、将列表 source 中的最后一个元素(从右侧)弹出，并返回给客户端。
// 2、将 source 弹出的元素插入（向左侧）到列表destination，作为destination列表的的头元素
// @param source string
// @param destination string
// @param timeout int64
func RPopLPush(source, destination string, timeout int64) string {
	if redisKey != "" {
		source = redisKey + ":" + source
		destination = redisKey + ":" + destination
	}

	return rdb.BRPopLPush(ctx, source, destination, time.Duration(timeout)*time.Second).Val()
}

// BRPopLPush RPopLPush的阻塞版本，当列表source为空时将阻塞连接，直到等待超时或有另一个客户端对source执行LPUSH或RPUSH命令为止
// @param source string
// @param destination string
// @param timeout int64
func BRPopLPush(source, destination string, timeout int64) string {
	if redisKey != "" {
		source = redisKey + ":" + source
		destination = redisKey + ":" + destination
	}

	return rdb.BRPopLPush(ctx, source, destination, time.Duration(timeout)*time.Second).Val()
}

// LIndex 通过索引获取列表中的元素
// @param key string
// @param index int64
// @return string
func LIndex(key string, index int64) string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LIndex(ctx, key, index).Val()
}

// LInsert 在列表的元素前或后插入元素
// @param key string
// @param where string before|after
// @param pivot string
// @param value string
// @return int64
// @return bool
func LInsert(key, where, pivot, value string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LInsert(ctx, key, where, pivot, value).Result()
}

// LLen 获取列表长度
// @param key string
// @return int64
func LLen(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LLen(ctx, key).Val()
}

// LRange 获取列表指定范围内的元素
// @param key string
// @param start int64
// @param stop int64
// @return []string
func LRange(key string, start, stop int64) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.LRange(ctx, key, start, stop).Val()
}

// LRem 根据参数count的值移除列表中与参数value相等的元素。count 的值可以是以下几种：
// 1、count > 0: 从表头开始向表尾搜索，移除与value相等的元素，数量为count
// 2、count < 0: 从表尾开始向表头搜索，移除与value相等的元素，数量为count的绝对值
// 3、count = 0: 移除表中所有与value相等的值
// @param key string
// @param count int64
// @return bool
func LRem(key string, count int64, value string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.LRem(ctx, key, count, value).Err()
	if err == nil {
		return true
	}
	return false
}

// LSet 设置指定下标的元素值
// @param key string
// @param index int64
// @param value string
// @return bool
func LSet(key string, index int64, value string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.LSet(ctx, key, index, value).Err()
	if err == nil {
		return true
	}
	return false
}

// SAdd 将一个或多个member元素加入到集合key当中，已经存在于集合的member元素将被忽略
// @param key string
// @param members []string
// @return int64
// @return error
func SAdd(key string, members []string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.SAdd(ctx, key, members).Result()
}

// SCard 返回集合key的基数(集合中元素的数量)
// @param key string
// @return int64
func SCard(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.SCard(ctx, key).Val()
}

// SDiff 返回一个集合的全部成员，该集合是所有给定集合之间的差集
// @param keys []string
// @return []string
func SDiff(keys []string) []string {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	return rdb.SDiff(ctx, keys...).Val()
}

// SDiffStore 与SDiff类似，但它将结果保存到destination集合
// 如果destination集合已经存在，则将其覆盖
// destination可以是key本身
// @param destination string
// @param keys []string
// @return bool
func SDiffStore(destination string, keys []string) bool {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	err := rdb.SDiffStore(ctx, destination, keys...).Err()

	if err == nil {
		return true
	}
	return false
}

// SInter 返回一个集合的全部成员，该集合是所有给定集合的交集
// @param keys []string
// @return []string
func SInter(keys []string) []string {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	return rdb.SInter(ctx, keys...).Val()
}

// SInterStore 与SInter类似，但它将结果保存到destination集合
// 如果destination集合已经存在，则将其覆盖
// destination可以是key本身
// @param destination string
// @param keys []string
// @return bool
func SInterStore(destination string, keys []string) bool {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	err := rdb.SInterStore(ctx, destination, keys...).Err()

	if err == nil {
		return true
	}
	return false
}

// SIsMember 判断member元素是否集合key的成员
// @param key string
// @param member string
// @return bool
func SIsMember(key string, member string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.SIsMember(ctx, key, member).Val()
}

// SMembers 返回集合 key 中的所有成员
// @param key string
// @return []string
func SMembers(key string) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.SMembers(ctx, key).Val()
}

// SMove 将member元素从source集合移动到destination集合
// @param key string
// @param destination string
// @return bool
func SMove(key, destination, member string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
		destination = redisKey + ":" + destination
	}

	return rdb.SMove(ctx, key, destination, member).Val()
}

// SRem 移除集合key中的一个或多个member元素，不存在的member元素会被忽略
// @param key string
// @param members []interface
// @return bool
func SRem(key string, members []interface{}) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.SRem(ctx, key, members...).Err()
	if err == nil {
		return true
	}
	return false
}

// SUnion 返回一个集合的全部成员，该集合是所有给定集合的并集
// @param keys []string
// @return []string
func SUnion(keys []string) []string {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	return rdb.SUnion(ctx, keys...).Val()
}

// SUnionStore 类似于SUnion命令，但它将结果保存到destination集合
// @param destination string
// @param keys []string
// @return bool
func SUnionStore(destination string, keys []string) bool {
	if redisKey != "" {
		for i := range keys {
			keys[i] = redisKey + ":" + keys[i]
		}
	}

	err := rdb.SUnionStore(ctx, destination, keys...).Err()
	if err == nil {
		return true
	}
	return false
}

// ZAdd 将一个或多个member元素及其score值加入到有序集key当中
// @param key string
// @param members map[interface{}]int64
// @return bool
func ZAdd(key string, members map[interface{}]int64) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.SAdd(ctx, key, members).Err()
	if err == nil {
		return true
	}
	return false
}

// ZCard 返回有序集key的基数
// @param key string
// @return int64
func ZCard(key string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZCard(ctx, key).Val()
}

// ZCount 返回有序集key中，score 值在min和max之间（默认包括score值等于min或max）的成员的数量
// @param key string
// @param min string
// @param max string
// @return int64
func ZCount(key, min, max string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZCount(ctx, key, min, max).Val()
}

// ZIncrBy 为有序集key的成员member的score值加上增量increment
// @param key string
// @param increment int64
// @param member string
// @return int64
// @return error
func ZIncrBy(key string, increment int64, member string) (int64, error) {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	result, err := rdb.ZIncrBy(ctx, key, float64(increment), member).Result()
	if err == nil {
		return int64(result), err
	}
	return 0, err
}

// ZRange 返回有序集key中，指定区间内的成员。
// 其中成员的位置按score值递增(从小到大)来排序。
// 具有相同score值的成员按字典序(lexicographical order )来排列
// @param key string
// @param start int64
// @param stop int64
// @return []string
func ZRange(key string, start, stop int64) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRange(ctx, key, start, stop).Val()
}

// ZRangeByScore 返回有序集ey中所有score 值介于min和max 之间(包括等于min或max)的成员。有序集成员按score值递增(从小到大)次序排列
// 具有相同 score 值的成员按字典序(lexicographical order)来排列(该属性是有序集提供的，不需要额外的计算)
// @param key string
// @param opt *redis.ZRangeBy
// @return []string
func ZRangeByScore(key string, opt *redis.ZRangeBy) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRangeByScore(ctx, key, opt).Val()
}

// ZRank 返回有序集 key 中成员 member 的排名。
// 其中有序集成员按 score 值递增(从小到大)顺序排列
// @param key string
// @param member string
// @return int64
func ZRank(key, member string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRank(ctx, key, member).Val()
}

// ZRem 移除有序集key中的一个或多个成员，不存在的成员将被忽略
// @param key string
// @param members []string
// @return bool
func ZRem(key string, members []string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.ZRem(ctx, key, members).Err()
	if err == nil {
		return true
	}
	return false
}

// ZRemRangeByRank 移除有序集key中指定排名(rank)区间内的所有成员
// @param key string
// @param opt *redis.ZRangeBy
// @return bool
func ZRemRangeByRank(key string, start, stop int64) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.ZRemRangeByRank(ctx, key, start, stop).Err()
	if err == nil {
		return true
	}
	return false
}

// ZRemRangeByScore 移除有序集key中指定分数（score）区间内的所有成员
// @param key string
// @param min string
// @param max string
// @return bool
func ZRemRangeByScore(key, min, max string) bool {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	err := rdb.ZRemRangeByScore(ctx, key, min, max).Err()
	if err == nil {
		return true
	}
	return false
}

// ZRevRange 返回有序集key中，指定区间内的成员。
// 其中成员的位置按score值递减(从大到小)来排列。
// 具有相同score值的成员按字典序的逆序(reverse lexicographical order)排列。
// @param key string
// @param start int64
// @param stop int64
// @return []string
func ZRevRange(key string, start, stop int64) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRange(ctx, key, start, stop).Val()
}

// ZRevRangeByLex 返回有序集key中指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列
// @param key string
// @param opt *redis.ZRangeBy
// @return []string
func ZRevRangeByLex(key string, opt *redis.ZRangeBy) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRangeByLex(ctx, key, opt).Val()
}

// ZRevRangeByScore 返回有序集key中指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列
// @param key string
// @param opt *redis.ZRangeBy
// @return []string
func ZRevRangeByScore(key string, opt *redis.ZRangeBy) []string {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRangeByScore(ctx, key, opt).Val()
}

// ZRevRangeByScoreWithScores 返回有序集key中指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列
// @param key string
// @param opt *redis.ZRangeBy
// @return []redis.Z
func ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) []redis.Z {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRangeByScoreWithScores(ctx, key, opt).Val()
}

// ZRevRangeWithScores 返回有序集key中指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列
// @param key string
// @param start int64
// @param stop int64
// @return []redis.Z
func ZRevRangeWithScores(key string, start, stop int64) []redis.Z {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRangeWithScores(ctx, key, start, stop).Val()
}

// ZRevRank 返回有序集key中成员member的排名。其中有序集成员按score值递减(从大到小)排序
// @param key string
// @param member string
// @return int64
func ZRevRank(key, member string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return rdb.ZRevRank(ctx, key, member).Val()
}

// ZScore 返回有序集key中成员member的score值
func ZScore(key, member string) int64 {
	if redisKey != "" {
		key = redisKey + ":" + key
	}

	return int64(rdb.ZScore(ctx, key, member).Val())
}
