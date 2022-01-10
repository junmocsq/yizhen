package comment

import (
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
)

type Comment struct {
	Id        uint32
	Uid       uint32 `gorm:"NOT NULL;comment:用户id"`
	MainId    uint32 `gorm:"NOT NULL;comment:评论所属模块主题id;index:main_idx,priority:1"`
	MainType  uint8  `gorm:"NOT NULL;default:1;comment:评论所属模块;index:main_idx,priority:2"`
	Content   string `gorm:"size:1000;NOT NULL;comment:评论内容"`
	ParentId  uint32 `gorm:"NOT NULL;default:0;comment:回复评论的一级评论id;index:main_idx,priority:4"`
	ParentUid uint32 `gorm:"NOT NULL;default:0;comment:回复评论的一级评论的uid"`
	ReplyId   uint32 `gorm:"NOT NULL;default:0;comment:被回复评论id"`
	ReplyUid  uint32 `gorm:"NOT NULL;default:0;comment:被回复评论uid"`
	CreatedAt int64  `gorm:"autoCreateTime;NOT NULL"`
	DeletedAt int64  `gorm:"NOT NULL;default:0;comment:删除时间;index:main_idx,priority:3"`
	Likely    uint32 `gorm:"NOT NULL;default:0;comment:评论喜欢数"`
	Unlikely  uint32 `gorm:"NOT NULL;default:0;comment:评论反对数"`
	ChildNum  uint32 `gorm:"NOT NULL;default:0;comment:子评论数"`
	ChildIds  string `gorm:"size:100;NOT NULL;default:'';comment:子评论ids"`
}

func (Comment) TableName() string {
	return "comments"
}

type comment struct {
}

func (c *comment) Tag() string {
	return "comment_comments"
}

func NewComment() *comment {
	return &comment{}
}

// Add 评论添加
func (c *comment) Add(uid, mainId, replyId uint32, mainType uint8, content string) int64 {
	com := Comment{
		Uid:       uid,
		MainId:    mainId,
		MainType:  mainType,
		Content:   content,
		ParentId:  0,
		ParentUid: 0,
		ReplyId:   replyId,
		ReplyUid:  0,
	}
	if replyId > 0 {
		replyComment := c.CommentById(replyId)
		if replyComment == nil {
			logrus.Errorf("comment find failed:%d", replyId)
			return 0
		}
		if replyComment.ParentId == 0 {
			com.ParentId = replyComment.Id
			com.ParentUid = replyComment.Uid
		} else {
			com.ParentId = replyComment.ParentId
			com.ParentUid = replyComment.ParentUid
		}
		com.ReplyId = replyComment.Id
		com.ReplyUid = replyComment.Uid
	}
	id := c.add(&com)
	if id > 0 {
		// todo 修改父评论child_ids
	}
	return id
}

func (c *comment) add(com *Comment) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Create(com).Statement
	n, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(com)
	if err != nil {
		logrus.Errorf("comment add err:%s", err)
		n = 0
	}
	return n
}

func (c *comment) CommentById(id uint32) *Comment {
	db := dbcache.NewDb()
	var com Comment
	stmt := db.DryRun().Find(&com, id).Statement
	err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&com)
	if err != nil {
		logrus.Errorf("comment find err:%s", err)
		return nil
	}
	return &com
}
