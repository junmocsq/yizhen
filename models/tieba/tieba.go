package tieba

import (
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Tieba 贴吧信息
type Tieba struct {
	Id        uint32 `json:"id"`
	Name      string `gorm:"size:20;NOT NULL;default:'';comment:贴吧名称;index:tb_idx,unique" json:"name"`
	Img       string `gorm:"size:150;NOT NULL;default:'';comment:贴吧头像" json:"img"`
	Desc      string `gorm:"size:200;NOT NULL;default:'';comment:贴吧介绍" json:"desc"`
	Uid       uint32 `gorm:"NOT NULL;default:0" json:"uid"`
	Status    uint8  `gorm:"NOT NULL;default:1;comment:状态 1 正常 2 审核 3 封禁" json:"status"`
	Follower  uint32 `gorm:"NOT NULL;default:0;comment:粉丝数" json:"follower"`
	Popu      uint32 `gorm:"NOT NULL;default:0;comment:人气值" json:"popu"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64  `gorm:"NOT NULL;default:0" json:"deleted_at"`
}

func (Tieba) TableName() string {
	return "tieba"
}

type tieba struct {
}

func NewTieba() *tieba {
	return &tieba{}
}
func (t *tieba) Tag() string {
	return "tieba_tieba"
}
func (t *tieba) Add(name, img, desc string, uid uint32) int64 {
	oldTieba := t.One(name)
	if oldTieba == nil || oldTieba.Id > 0 {
		logrus.Errorf("tieba:%s 已存在", name)
		return 0
	}
	db := dbcache.NewDb()
	var tb Tieba
	tb.Name = name
	tb.Img = img
	tb.Desc = desc
	tb.Uid = uid
	stmt := db.DryRun().Create(&tb).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&tb)
	if err != nil {
		logrus.Errorf("tieba add err:%s", err)
		n = 0
	}
	return n
}

func (t *tieba) One(name string) *Tieba {
	db := dbcache.NewDb()
	var tb Tieba
	stmt := db.DryRun().Find(&tb, "name = ?", name).Statement
	err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&tb)

	if err != nil {
		logrus.Errorf("tieba find err:%s", err)
		return nil
	}
	return &tb
}

func (t *tieba) UpdateInfo(id uint32, img, desc string) int64 {
	db := dbcache.NewDb()

	stmt := db.DryRun().Model(&Tieba{}).Select("img", "desc").
		Where("id = ?", id).Updates(Tieba{Img: img, Desc: desc}).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()

	if err != nil {
		logrus.Errorf("tieba update err:%s", err)
		return 0
	}
	return n
}

func (t *tieba) UpdateStatus(id uint32, status uint8) int64 {
	db := dbcache.NewDb()

	stmt := db.DryRun().Model(&Tieba{}).
		Where("id = ?", id).Update("status", status).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba update err:%s", err)
		return 0
	}
	return n
}

func (t *tieba) UpdateFollower(id uint32, follower uint32) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(&Tieba{}).Where("id = ?", id).Omit("UpdatedAt").
		Update("follower", follower).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba update follower:%s", err)
		return 0
	}
	return n
}

func (t *tieba) IncrFollower(id uint32, follower uint32) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(&Tieba{}).Where("id = ?", id).Omit("UpdatedAt").
		Update("follower", gorm.Expr("follower + ?", follower)).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba update follower:%s", err)
		return 0
	}
	return n
}

func (t *tieba) UpdatePopu(id uint32, popu uint32) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(&Tieba{}).Where("id = ?", id).Omit("UpdatedAt").
		Update("popu", popu).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba update popu:%s", err)
		return 0
	}
	return n
}

func (t *tieba) IncrPopu(id uint32, popu uint32) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(&Tieba{}).Where("id = ?", id).Omit("UpdatedAt").
		Update("popu", gorm.Expr("popu + ?", popu)).Statement
	n, err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("tieba incr popu:%s", err)
		return 0
	}
	return n
}

func (t *tieba) List(page int, size int) []Tieba {
	db := dbcache.NewDb()
	var tbs []Tieba
	stmt := db.DryRun().Where("status = 1").Order("popu DESC,id DESC").Limit(size).Offset((page - 1) * size).
		Find(&tbs).Statement
	err := db.SetTag(t.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&tbs)
	if err != nil {
		logrus.Errorf("tieba incr popu:%s", err)
		return nil
	}
	return tbs
}
