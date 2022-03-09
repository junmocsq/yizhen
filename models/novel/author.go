package novel

type Author struct {
	Id        int    `json:"id"`
	Nick      string `json:"nick"`
	WhatsUp   string `json:"whats_up"`
	Avatar    string `json:"avatar"`
	Followers int    `json:"followers"`
	Novels    int    `json:"novels"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64  `gorm:"NOT NULL;default:0" json:"deleted_at"`
}

func (Author) TableName() string {
	return "nv_author"
}
