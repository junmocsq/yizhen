package tieba

// Manager 贴吧管理人员
type Manager struct {
	Id        uint32 `json:"id"`
	Uid       uint32 `json:"uid"`
	Tid       uint32 `gorm:"NOT NULL;comment:贴吧ID;index:tid_idx" json:"tid"`
	Role      uint8  `gorm:"NOT NULL;default 2;comment:角色 1 吧主 2 小吧主" json:"role"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
}

func (Manager) TableName() string {
	return "tieba_manager"
}
