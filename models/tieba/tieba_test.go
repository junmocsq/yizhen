package tieba

import (
	"testing"
)

func TestTieba_Add(t *testing.T) {
	NewTieba().Add("赵小凡", "", "赵小凡在干嘛", 1)
	NewTieba().Add("澄清杨", "", "澄清杨想吃小龙虾", 5)
	NewTieba().UpdateInfo(1, "", "赵小凡在干嘛呢")
	NewTieba().UpdateStatus(1, 2)
	NewTieba().UpdateStatus(1, 1)

	NewTieba().UpdateFollower(1, 200)
	NewTieba().IncrFollower(1, 10)

	NewTieba().UpdatePopu(1, 300)
	NewTieba().IncrPopu(1, 20)
	tb := NewTieba().One("赵小凡")
	if tb.Follower != 210 || tb.Popu != 320 {
		t.Error("tieba test failed")
	}
	if len(NewTieba().List(1, 10)) < 2 {
		t.Error("tieba list find failed")
	}
}
