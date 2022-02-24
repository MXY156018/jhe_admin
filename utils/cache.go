package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var goCache *cache.Cache

func init() {
	goCache = cache.New(0, 0)
}

//设置一个cache  可以用于增加和更新
func SetCache(k string, v interface{}, t time.Duration) {
	goCache.Set(k, v, t)
}

//新增cache 只有当k不存在或已过期时可以增加成功  否则返回error
func AddCache(k string, v interface{}, t time.Duration) error {
	err := goCache.Add(k, v, t)
	return err
}

//跟set用法一样  只是使用默认的有效时间
func SetDefault(k string, v interface{}) {
	goCache.SetDefault(k, v)

}

//按k获取  如果cache中没有key，返回的value为nil，同时返回一个bool类型的参数表示key是否存在。
func GetCache(k string) (interface{}, bool) {
	v, b := goCache.Get(k)
	return v, b
}

//与Get接口的区别是，返回参数中增加了key有效期的信息，如果是不会过期的key，返回的是time.Time类型的零值。
func GetWithExpiration(k string) (interface{}, time.Time, bool) {
	v, t, b := goCache.GetWithExpiration(k)
	return v, t, b
}

//如果key存在且为过期，将对应value更新为新的值；否则返回error。
func Replace(k string, x interface{}, d time.Duration) error {
	err := goCache.Replace(k, x, d)
	return err
}

//对于cache中value是int, int8, int16, int32, int64, uintptr, uint,uint8, uint32, or uint64, float32,float64这些类型记录，
//可以使用该接口，将value值减n。如果key不存在或value不是上述类型，会返回error。
func Decrement(k string, n int64) error {
	err := goCache.Decrement(k, n)
	return err
}

//DecrementXXX:对于Decrement接口中提到的各种类型，还有对应的接口来处理，同时这些接口可以得到value变化后的结果。
//返回value-n后的结果
func DecrementInt64(k string, n int64) (int64, error) {
	v, err := goCache.DecrementInt64(k, n)
	return v, err
}

//同Decrement一样用法
//可以使用该接口，将value值加n。如果key不存在或value不是上述类型，会返回error。
func Increment(k string, n int64) error {
	err := goCache.Increment(k, n)
	return err
}

//同DecrementXXX一样 返回value + n后的值
func IncrementInt64(k string, n int64) (int64, error) {
	v, err := goCache.IncrementInt64(k, n)
	return v, err
}

//删除一个k-v
func Delete(k string) {
	goCache.Delete(k)
}

//删除所有已过期的k-v
func DeleteExpired() {
	goCache.DeleteExpired()
}

//清空所有k-v
func Flush() {
	goCache.Flush()
}

//返回cache中的记录数量。
//需要注意的是，返回的数值可能会比实际能获取到的数值大，对于已经过期但还没有及时清理的记录也会被统计。
func ItemCount() int {
	return goCache.ItemCount()
}

func OnEvicted(f func(string, interface{})) {
	goCache.OnEvicted(f)
}
func f(k string, v interface{}) (string, interface{}) {
	return k, v
}
