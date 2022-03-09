package tieba

import "testing"

func TestFollow_CheckFollow(t *testing.T) {
	var uid, tid uint32
	uid = 100
	tid = 1

	f := NewFollow()
	t.Log(f.Delete(tid, uid))
	t.Log(f.CheckFollow(tid, uid))

	//t.Log(f.Add(tid,uid))
	//t.Log(f.CheckFollow(tid,uid))
}
