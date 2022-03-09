package novel

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAuthorFollow(t *testing.T) {
	f := NewAuthorFollow()
	for uid := 10; uid < 20; uid++ {
		for i := 100; i > 90; i-- {
			aid := i - uid
			if !f.CheckFollow(aid, uid) {
				f.Add(aid, uid)
			}
		}
	}
	uid := 999
	aid := 999
	f.Delete(aid, uid)
	Convey("Author follow", t, func() {
		Convey("Add", func() {
			So(f.Add(aid, uid), ShouldBeGreaterThan, 0)
		})
		Convey("Delete", func() {
			So(f.Delete(aid, uid), ShouldBeGreaterThan, 0)
		})
		Convey("Followers", func() {
			size := f.FollowerNum(90)
			So(len(f.Followers(90, 1, size)), ShouldEqual, size)
		})
		Convey("Followings", func() {
			size := f.FollowingNum(10)
			So(len(f.Followings(10, 1, size)), ShouldEqual, size)
		})
	})
}
