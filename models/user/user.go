package user

type User struct {
	Id        uint32 `json:"id"`
	Avatar    string `gorm:"size:150;NOT NULL;default:'';comment:头像" json:"avatar"`
	Name      string `gorm:"size:50;NOT NULL;default:'';comment:昵称" json:"name"`
	Sex       uint8  `gorm:"NOT NULL;default:0;comment:0 未知 1 男 2 女" json:"sex"`
	Email     string `gorm:"size:100;NOT NULL;default:'';comment:邮箱" json:"email"`
	Phone     string `gorm:"size:20;NOT NULL;default:'';comment:电话" json:"phone"`
	Nation    int32  `gorm:"NOT NULL;default:86;comment:电话区号" json:"nation"`
	WhatIsUp  string `gorm:"size:100;NOT NULL;default:'';comment:心情" json:"what_is_up"`
	Country   string `gorm:"size:20;NOT NULL;default:'';comment:国家" json:"country"`
	Province  string `gorm:"size:20;NOT NULL;default:'';comment:省" json:"province"`
	City      string `gorm:"size:20;NOT NULL;default:'';comment:市" json:"city"`
	CreatedAt int64  `gorm:"autoCreateTime;comment:注册时间" json:"created_at"`
	UpdatedAt int64  `gorm:"NOT NULL;default:0;comment:信息更新时间" json:"updated_at"`
	DeletedAt int64  `gorm:"NOT NULL;default:0;comment:删除时间" json:"deleted_at"`
	LoginAt   int64  `gorm:"NOT NULL;default:0;comment:最后登录时间" json:"login_at"`
	Status    uint8  `gorm:"NOT NULL;default:1;comment:状态 1 正常 2 审核中 3 永久封禁" json:"status"`
	Passwd    string `gorm:"size:40;NOT NULL;default:'';comment:密码" json:"passwd"`
	QQ        string `gorm:"size:100;NOT NULL;default:'';comment:qq登录openid" json:"qq"`
	Wechat    string `gorm:"size:100;NOT NULL;default:'';comment:微信登录openid" json:"wechat"`
	Weibo     string `gorm:"size:100;NOT NULL;default:'';comment:微博登录openid" json:"weibo"`
	Apple     string `gorm:"size:40;NOT NULL;default:'';comment:苹果登录openid" json:"apple"`
	Gold      int64  `gorm:"comment:金币" json:"gold"`
}
