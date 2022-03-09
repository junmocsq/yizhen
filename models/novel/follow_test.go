package novel

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNovelFollow(t *testing.T) {
	f := NewFollow()
	for uid := 10; uid < 20; uid++ {
		for i := 100; i > 90; i-- {
			nid := i - uid
			if !f.CheckFollow(nid, uid) {
				f.Add(nid, uid)
			}
		}
	}
	uid := 999
	nid := 999
	f.Delete(nid, uid)
	Convey("Novel follow", t, func() {
		Convey("Add", func() {
			So(f.Add(nid, uid), ShouldBeGreaterThan, 0)
		})
		Convey("Delete", func() {
			So(f.Delete(nid, uid), ShouldBeGreaterThan, 0)
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
