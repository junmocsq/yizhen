package novel

type Novel struct {
	Id                int
	Name              string
	Img               string
	Intro             string
	SimpleIntro       string
	Aid               int
	Category          int
	Status            uint8 `json:"status" gorm:"comment:1 正常，2 审核，3 关闭"`
	State             uint8 `json:"state" gorm:"comment:1 连载，2 完结，3 断更"`
	Words             int
	Chapters          int
	LastChapterId     int
	LastChapterName   string
	LastChapterUpdate int64

	FinishedAt int64 `gorm:"comment:完结时间" json:"finished_at"`
	UpdatedAt  int64 `gorm:"comment:更新时间" json:"updated_at"`
	CreatedAt  int64 `gorm:"autoCreateTime" json:"created_at"`
}
