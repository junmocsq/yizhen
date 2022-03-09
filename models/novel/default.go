package novel

import (
	"github.com/junmocsq/yizhen/models/dbcache"
)

func init() {
	db := dbcache.NewDb()
	db.DB().AutoMigrate(&Category{}, &Follow{}, &AuthorFollow{})
}
