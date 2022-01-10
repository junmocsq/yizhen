package comment

type Notify struct {
	Id        uint64
	Uid       uint32 `gorm:"NOT NULL;default:0;comment:通知所属UID"`
	MainId    uint32 `gorm:"NOT NULL;default:0;comment:模块内容id"`
	MainType  uint8  `gorm:"NOT NULL;default:1;comment:模块"`
	CUid      uint32 `gorm:"NOT NULL;default:0;comment:一级评论UID"`
	CCid      uint32 `gorm:"NOT NULL;default:0;comment:一级评论ID"`
	CContent  string `gorm:"size:1000;NOT NULL;default:0;comment:一级评论内容"`
	RUid      uint32 `gorm:"NOT NULL;default:0;comment:回复评论UID"`
	RCid      uint32 `gorm:"NOT NULL;default:0;comment:回复评论ID"`
	RRUid     uint32 `gorm:"NOT NULL;default:0;comment:被回复评论UID"`
	RContent  string `gorm:"size:1000;NOT NULL;default:0;comment:回复评论内容"`
	CreatedAt int64  `gorm:"autoCreateTime;NOT NULL"`
	DeletedAt int64  `gorm:"NOT NULL;default:0;comment:删除时间"`
}

func (Notify) TableName() string {
	return "comment_notifies"
}
