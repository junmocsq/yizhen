package novel

// 用户划线内容，包含计数

type UnderlineContent struct {
	Id        int    `json:"id"`
	Nid       int    `json:"nid"`
	Cid       int    `json:"cid"`
	IsValid   uint8  `json:"is_valid"  gorm:"comment:0 无效（由于修改内容导致丢失等） 1 有效;NOT NULL;default:1"`
	Users     int    `json:"users"  gorm:"comment:划线用户数"`
	Start     int    `json:"start" gorm:"comment:标记起始"`
	End       int    `json:"end" gorm:"comment:标记结束"`
	Md5Hash   string `json:"md5_hash" gorm:"comment:标记内容hash值;size:32"`
	Content   string `json:"content"  gorm:"comment:标记内容;size:3000"`
	CreatedAt int64  `json:"created_at" gorm:"comment:autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UnderlineContent) TableName() string {
	return "nv_underline_content"
}
