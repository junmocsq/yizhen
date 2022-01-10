package dbcache

type Cache interface {
	Set(key, value string, expire ...int) bool
	Get(key string) string
	Expire(key string, expire int) int
	Del(key string) int
}

// RegisterCache 缓存注册
func RegisterCache(c Cache) {
	cacheHandler = c
}
