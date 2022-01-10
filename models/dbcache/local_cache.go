package dbcache

// 本地缓存
import (
	"github.com/patrickmn/go-cache"
	"time"
)

var localCacheHandler *cache.Cache

var _ Cache = (*localCache)(nil)

type localCache struct {
}

func NewLocalCache() *localCache {
	if localCacheHandler == nil {
		localCacheHandler = cache.New(5*time.Minute, 10*time.Minute)
	}
	return &localCache{}
}
func (r *localCache) Set(key, value string, expire ...int) bool {
	if len(expire) > 0 {
		localCacheHandler.Set(key, value, time.Duration(expire[0])*time.Second)
	} else {
		localCacheHandler.Set(key, value, cache.NoExpiration)
	}
	return true
}
func (r *localCache) Get(key string) string {
	res, ok := localCacheHandler.Get(key)
	if !ok {
		return ""
	}
	return res.(string)
}
func (r *localCache) Del(key string) int {
	localCacheHandler.Delete(key)
	return 1
}
func (r *localCache) Expire(key string, expire int) int {
	value := r.Get(key)
	if value == "" {
		return 0
	}
	r.Set(key, value, expire)
	return 1
}
