package novel

import (
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Follow struct {
	Uid       int   `gorm:"NOT NULL;comment:UID;index:uid_idx,unique,priority:1" json:"uid"`
	Nid       int   `gorm:"NOT NULL;comment:小说id;index:uid_idx,unique,priority:2;index:nid_sort_idx,priority:1" json:"nid"`
	CreatedAt int64 `gorm:"autoCreateTime;index:nid_sort_idx,priority:2" json:"created_at"`
}

func (Follow) TableName() string {
	return "nv_follow"
}
func (f *follow) Tag(uid int) string {
	return "nv_follow_" + strconv.Itoa(uid)
}

type follow struct {
}

func NewFollow() *follow {
	return &follow{}
}

func (f *follow) CheckFollow(nid, uid int) bool {
	db := dbcache.NewDb()
	var fl Follow
	stmt := db.DryRun().Where("nid=? AND uid=?", nid, uid).Find(&fl).Statement
	err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&fl)

	if err != nil {
		logrus.Errorf("novel follow find err:%s", err)
		return false
	}
	return fl.Uid > 0
}

func (f *follow) Add(nid, uid int) int {
	db := dbcache.NewDb()
	var fl = Follow{
		Uid: uid,
		Nid: nid,
	}
	stmt := db.DryRun().Create(&fl).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("novel follow add err:%s", err)
		return 0
	}
	return int(n)
}

func (f *follow) Delete(nid, uid int) int {
	db := dbcache.NewDb()
	stmt := db.DryRun().Where("nid=? AND uid=?", nid, uid).Delete(&Follow{}).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("novel follow delete err:%s", err)
		return 0
	}
	return int(n)
}

// 作品的被收藏记录
func (f *follow) Followers(nid, page, size int) []Follow {
	db := dbcache.NewDb()
	var followers []Follow
	stmt := db.DryRun().Order("created_at desc").Where("nid=?", nid).Limit(size).Offset(page*size - size).Find(&followers).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followers)
	if err != nil {
		logrus.Errorf("find novel followers list err:%s", err)
		return nil
	}
	return followers
}

func (f *follow) FollowerNum(nid int) int {
	db := dbcache.NewDb()
	var followerNum int64
	stmt := db.DryRun().Model(&Follow{}).Select("COUNT(*) num").Where("nid=?", nid).Find(&followerNum).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followerNum)
	if err != nil {
		logrus.Errorf("find novel followers num err:%s", err)
		return 0
	}
	return int(followerNum)
}

// 用户的收藏作品记录
func (f *follow) Followings(uid, page, size int) []Follow {
	db := dbcache.NewDb()
	var followers []Follow
	stmt := db.DryRun().Order("created_at desc").Where("uid=?", uid).Limit(size).Offset(page*size - size).Find(&followers).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followers)
	if err != nil {
		logrus.Errorf("find user followings novel list err:%s", err)
		return nil
	}
	return followers
}

func (f *follow) FollowingNum(uid int) int {
	db := dbcache.NewDb()
	var followingNum int64
	stmt := db.DryRun().Model(&Follow{}).Select("COUNT(*) num").Where("uid=?", uid).Find(&followingNum).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followingNum)
	if err != nil {
		logrus.Errorf("find user followings novel num err:%s", err)
		return 0
	}
	return int(followingNum)
}
