package comment

type Like struct {
	Id     uint32
	Uid    uint32 `gorm:"NOT NULL;comment:UID"`
	IsLike int8   `gorm:"NOT NULL;default:1;comment:-1 不喜欢，1 喜欢"`
}

func (Like) TableName() string {
	return "comment_likes"
}
