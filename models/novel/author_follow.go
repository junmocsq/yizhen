package novel

import (
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
	"strconv"
)

type AuthorFollow struct {
	Uid       int   `gorm:"NOT NULL;comment:UID;index:uid_idx,unique,priority:1" json:"uid"`
	Aid       int   `gorm:"NOT NULL;comment:作者id;index:uid_idx,unique,priority:2;index:aid_sort_idx,priority:1" json:"aid"`
	CreatedAt int64 `gorm:"autoCreateTime;index:aid_sort_idx,priority:1" json:"created_at"`
}

func (AuthorFollow) TableName() string {
	return "nv_author_follow"
}
func (f *authorFollow) Tag(uid int) string {
	return "nv_author_follow_" + strconv.Itoa(uid)
}

type authorFollow struct {
}

func NewAuthorFollow() *authorFollow {
	return &authorFollow{}
}

func (f *authorFollow) CheckFollow(aid, uid int) bool {
	db := dbcache.NewDb()
	var fl AuthorFollow
	stmt := db.DryRun().Where("aid=? AND uid=?", aid, uid).Find(&fl).Statement
	err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&fl)

	if err != nil {
		logrus.Errorf("author follow find err:%s", err)
		return false
	}
	return fl.Uid > 0
}

func (f *authorFollow) Add(aid, uid int) int {
	db := dbcache.NewDb()
	var fl = AuthorFollow{
		Uid: uid,
		Aid: aid,
	}
	stmt := db.DryRun().Create(&fl).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("author follow add err:%s", err)
		return 0
	}
	return int(n)
}

func (f *authorFollow) Delete(aid, uid int) int {
	db := dbcache.NewDb()
	stmt := db.DryRun().Where("aid=? AND uid=?", aid, uid).Delete(&AuthorFollow{}).Statement
	n, err := db.SetTag(f.Tag(uid)).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("author follow delete err:%s", err)
		return 0
	}
	return int(n)
}

// Followers 作品的被收藏记录
func (f *authorFollow) Followers(aid, page, size int) []AuthorFollow {
	db := dbcache.NewDb()
	var followers []AuthorFollow
	stmt := db.DryRun().Order("created_at desc").Where("aid=?", aid).Limit(size).Offset(page*size - size).Find(&followers).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followers)
	if err != nil {
		logrus.Errorf("find author followers list err:%s", err)
		return nil
	}
	return followers
}

func (f *authorFollow) FollowerNum(aid int) int {
	db := dbcache.NewDb()
	var followerNum int64
	stmt := db.DryRun().Model(&AuthorFollow{}).Select("COUNT(*) num").Where("aid=?", aid).Find(&followerNum).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followerNum)
	if err != nil {
		logrus.Errorf("find author followers num err:%s", err)
		return 0
	}
	return int(followerNum)
}

// 用户的收藏作者记录
func (f *authorFollow) Followings(uid, page, size int) []AuthorFollow {
	db := dbcache.NewDb()
	var followers []AuthorFollow
	stmt := db.DryRun().Order("created_at desc").Where("uid=?", uid).Limit(size).Offset(page*size - size).Find(&followers).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followers)
	if err != nil {
		logrus.Errorf("find user followings author list err:%s", err)
		return nil
	}
	return followers
}

func (f *authorFollow) FollowingNum(uid int) int {
	db := dbcache.NewDb()
	var followingNum int64
	stmt := db.DryRun().Model(&AuthorFollow{}).Select("COUNT(*) num").Where("uid=?", uid).Find(&followingNum).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&followingNum)
	if err != nil {
		logrus.Errorf("find user followings author num err:%s", err)
		return 0
	}
	return int(followingNum)
}
