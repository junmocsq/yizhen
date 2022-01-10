package tieba

// Moment 贴吧动态
type Moment struct {
	Id        uint32 `json:"id"`
	Tid       uint32 `json:"tid"`
	Uid       uint32 `json:"uid"`
	Content   string `gorm:"size:2000;NOT NULL;comment:帖子内容" json:"content"`          // 2000
	Images    string `gorm:"size:1500;NOT NULL;comment:帖子图片" json:"images"`           // 1500
	Status    uint8  `gorm:"NOT NULL;default:1;comment:1 正常 2 审核 3 封禁" json:"status"` // 1
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64  `gorm:"NOT NULL;default:0" json:"deleted_at"`
}

func (Moment) TableName() string {
	return "tieba_moment"
}
