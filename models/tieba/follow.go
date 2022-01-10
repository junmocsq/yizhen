package tieba

// Follow 贴吧关注者
type Follow struct {
	Uid       uint32 `gorm:"NOT NULL;comment:UID;index:uid_idx,unique,priority:1" json:"uid"`
	Tid       uint32 `gorm:"NOT NULL;comment:贴吧ID;index:uid_idx,unique,priority:2" json:"tid"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (Follow) TableName() string {
	return "tieba_follow"
}
