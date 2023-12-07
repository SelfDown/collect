package collect

import (
	utils "collect.mod/src/collect/utils"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"strings"
	"time"
)

/**
处理缓存
*/
const CacheGetName = "GET"            // 单个获取
const CacheSetName = "SET"            // 单个设置
const BulkSetCache = "BULK_SET_CACHE" // 批量设置

var localCache *ristretto.Cache

type CacheHandler struct {
}

func init() {
	// 初始内存缓存连接
	getCacheObj()

}
func getCacheObj() *ristretto.Cache {
	if localCache != nil {
		return localCache
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil
	}
	localCache = cache
	return localCache

}
func (t *CacheHandler) GetCacheKey(room string, fields []string, params map[string]interface{}) string {
	valueList := utils.GetFieldValueList(fields, params)
	valueStr := strings.Join(valueList, ":")
	return fmt.Sprintf("%s[%s]", room, valueStr)
}

// Get 获取缓存
func (t *CacheHandler) Get(key interface{}) (interface{}, bool) {
	cache := getCacheObj()
	return cache.Get(key)
}

// Set 设置缓存
func (t *CacheHandler) Set(key, value interface{}, second int64) bool {
	cache := getCacheObj()
	ok := cache.SetWithTTL(key, value, 0, time.Duration(second)*time.Second)
	return ok
}

// SetMini 设置缓存,毫秒级别
func (t *CacheHandler) SetMini(key, value interface{}, second int64) bool {
	cache := getCacheObj()
	ok := cache.SetWithTTL(key, value, 0, time.Duration(second)*time.Millisecond)

	return ok
}

// GetTTl 设置缓存,毫秒级别
func (t *CacheHandler) GetTTl(key interface{}) (time.Duration, bool) {
	cache := getCacheObj()
	return cache.GetTTL(key)
}

// Wait 设置缓存
func (t *CacheHandler) Wait() {
	cache := getCacheObj()
	cache.Wait()
}
