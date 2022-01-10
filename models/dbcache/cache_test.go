package dbcache

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestLocalCache(t *testing.T) {
	c := NewLocalCache()
	RegisterCache(c)
	arr := map[string]string{
		"junmo":      "csq",
		"caohongdou": "菜红豆",
	}
	Convey("local cache", t, func() {
		Convey("set key value", func() {
			for k, v := range arr {
				c.Set(k, v, 300)
			}
		})
		Convey("get value", func() {
			for k, v := range arr {
				So(c.Get(k), ShouldEqual, v)
			}
		})
		Convey("del value ,value is empty", func() {
			for k, _ := range arr {
				c.Del(k)
				So(c.Get(k), ShouldBeEmpty)
			}
		})
		Convey("set expire", func() {
			for k, v := range arr {
				c.Set(k, v, 300)
				vv, tt, ok := localCacheHandler.GetWithExpiration(k)
				So(ok, ShouldBeTrue)
				So(vv.(string), ShouldEqual, v)
				So(tt.Unix()-time.Now().Unix(), ShouldBeBetween, 298, 301)
				c.Expire(k, 5)
				vv, tt, ok = localCacheHandler.GetWithExpiration(k)
				So(ok, ShouldBeTrue)
				So(vv.(string), ShouldEqual, v)
				So(tt.Unix()-time.Now().Unix(), ShouldBeBetween, 3, 6)
				c.Del(k)
			}
		})
	})
}
func TestRedisCache(t *testing.T) {
	RedisCacheInit("127.0.0.1", "6379", "")
	c := NewRedisCache()
	RegisterCache(c)
	arr := map[string]string{
		"junmo":      "csq",
		"caohongdou": "菜红豆",
	}
	Convey("redisCacheHandler cache", t, func() {
		Convey("set key value", func() {
			for k, v := range arr {
				c.Set(k, v, 300)
			}
		})
		Convey("get value", func() {
			for k, v := range arr {
				So(c.Get(k), ShouldEqual, v)
			}
		})
		Convey("del value ,value is empty", func() {
			for k, _ := range arr {
				c.Del(k)
				So(c.Get(k), ShouldBeEmpty)
			}
		})
		Convey("set expire", func() {
			for k, v := range arr {
				c.Set(k, v, 300)
				So(redisCacheHandler.TTL(k), ShouldBeBetween, 298, 301)
				c.Set(k, v, 5)
				So(redisCacheHandler.TTL(k), ShouldBeBetween, 4, 6)
				c.Del(k)
			}
		})
	})
}
