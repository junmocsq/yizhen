package dbcache

import "github.com/junmocsq/jlib/jredis"

var (
	redisModule       = "sql"
	redisCacheHandler = jredis.NewRedis(redisModule)
	isRegisterRedis   = false
)

func SetDebug(debug ...bool) {
	if len(debug) > 0 {
		jredis.SetDebug(debug[0])
	}
}
func RedisCacheInit(host, port, auth string, module ...string) {
	if isRegisterRedis {
		return
	}
	if len(module) > 0 {
		redisModule = module[0]
	}
	jredis.RegisterRedisPool(host, port, jredis.ModuleConf(redisModule), jredis.AuthConf(auth), jredis.PrefixConf(redisModule))
	cacheHandler = NewRedisCache()
	isRegisterRedis = true
}

func NewRedisCache() *redisCache {
	return &redisCache{}
}

var _ Cache = (*redisCache)(nil)

type redisCache struct {
}

func (r *redisCache) Set(key, value string, expire ...int) bool {
	if len(expire) > 0 {
		return redisCacheHandler.SETEX(key, value, expire[0])
	} else {
		return redisCacheHandler.SET(key, value)
	}
}
func (r *redisCache) Get(key string) string {
	return redisCacheHandler.GET(key)
}
func (r *redisCache) Del(key string) int {
	return redisCacheHandler.DEL(key)
}
func (r *redisCache) Expire(key string, expire int) int {
	return redisCacheHandler.EXPIRE(key, expire)
}
