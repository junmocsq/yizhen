package tieba

import (
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
	"strconv"
)

// Follow 贴吧关注者
type Follow struct {
	Uid       uint32 `gorm:"NOT NULL;comment:UID;index:uid_idx,unique,priority:1" json:"uid"`
	Tid       uint32 `gorm:"NOT NULL;comment:贴吧ID;index:uid_idx,unique,priority:2" json:"tid"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (Follow) TableName() string {
	return "tieba_follow"
}
func (f *follow) Tag(uid uint32) string {
	return "tieba_follow_" + strconv.Itoa(int(uid))
}

type follow struct {
}

func NewFollow() *follow {
	return &follow{}
}

func (f *follow) CheckFollow(tid, uid uint32) bool {
	db := dbcache.NewDb()
	var fl Follow
	stmt := db.DryRun().Where("tid=? AND uid=?", tid, uid).Find(&fl).Statement
	err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&fl)

	if err != nil {
		logrus.Errorf("tieba follow find err:%s", err)
		return false
	}
	return fl.Uid > 0
}

func (f *follow) Add(tid, uid uint32) int64 {
	db := dbcache.NewDb()
	var fl = Follow{
		Uid: uid,
		Tid: tid,
	}
	stmt := db.DryRun().Create(&fl).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba follow add err:%s", err)
		return 0
	}
	return n
}

func (f *follow) Delete(tid, uid uint32) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Where("tid=? AND uid=?", tid, uid).Delete(&Follow{}).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba follow delete err:%s", err)
		return 0
	}
	return n
}

func (f *follow) List(tid uint32, page, size int) []Follow {
	db := dbcache.NewDb()
	var followers []Follow
	stmt := db.DryRun().Order("created_at desc").Where("tid=?", tid).Limit(size).Offset(page*size - size).Find(&followers).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followers)
	if err != nil {
		logrus.Errorf("tieba follow find list err:%s", err)
		return nil
	}
	return followers
}
