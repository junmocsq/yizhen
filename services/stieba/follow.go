package stieba

import (
	"fmt"
	"github.com/junmocsq/yizhen/models/tieba"
)

type follow struct {
}

func NewSFollow() *follow {
	return &follow{}
}

// Follow 0 操作失败 1 关注成功 2 取消关注
func (f *follow) Follow(uid, tid uint32) (result int) {
	var follower int32
	defer func() {
		if follower != 0 {
			tieba.NewTieba().IncrFollower(tid, follower)
		}
	}()
	fm := tieba.NewFollow()
	res := fm.CheckFollow(tid, uid)
	if res {
		n := fm.Delete(tid, uid)
		if n == 0 {
			return
		} else {
			follower = -1
			return 2
		}
	} else {
		n := fm.Add(tid, uid)
		if n == 0 {
			return
		} else {
			follower = 1
			return 1
		}
	}
}

func (f *follow) Followers(tid uint32, page, size int) {
	res := tieba.NewFollow().List(tid, page, size)
	fmt.Println(res)
}
