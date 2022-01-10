package tieba

import (
	"github.com/junmocsq/yizhen/models/dbcache"
)

func init() {
	db := dbcache.NewDb()
	db.DB().AutoMigrate(&Tieba{}, &Follow{}, &Manager{}, &Moment{})
}
