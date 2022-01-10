package dbcache

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

type User struct {
	Id         int    `gorm:"primaryKey;type:int;NOT NULL"`
	Nickname   string `gorm:"uniqueIndex;size:50;NOT NULL;default:''"`
	Avatar     string `gorm:"size:50;NOT NULL;default:''"`
	Passwd     string `gorm:"size:40;NOT NULL;default:''"`
	Email      string `gorm:"size:50;NOT NULL;default:''"`
	Phone      int64  `gorm:"type:int;NOT NULL;default:0"`
	NationCode int    `gorm:"type:int;NOT NULL;default:86"`
	Salt       string `gorm:"type:char;size:6;NOT NULL;default:''"`
	Gender     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default:0;comment:0 unknown 1 male 2 female"`
	WhatIsUp   string `gorm:"size:100;NOT NULL;default:''"`
	LoginTime  int64  `gorm:"NOT NULL;default:0"`
	UpdatedAt  int64  `gorm:"autoUpdateTime;NOT NULL;default:0"`
	CreatedAt  int64  `gorm:"autoCreateTime;NOT NULL;default:0"`
}

func (User) TableName() string {
	return "user_test"
}
func TestCache(t *testing.T) {
	db := NewDb()
	db.DB().AutoMigrate(&User{})
}

func ExampleDao_Exec() {
	db := NewDb()
	var user1 = User{
		Id:       1,
		Nickname: "junmocsq1",
		Email:    "",
	}
	db.DB().AutoMigrate(&User{})

	tag := "users000"
	stmt := db.DryRun().Delete(&user1).Statement
	db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	stmt = db.DryRun().Create(&user1).Statement
	n, err := db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	fmt.Println(n, err)
	// OutPut: 1 <nil>
}

func ExampleDao_Fetch() {
	db := NewDb()
	var user1 User
	db.DB().AutoMigrate(&User{})
	stmt := db.DryRun().Find(&user1, 1).Statement
	tag := "users000"
	err := db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user1)
	fmt.Println(user1.Id, err)

	var users []User
	// 会话模式
	stmt = db.DryRun().Find(&users).Statement
	db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&users)
	fmt.Println("result", len(users) > 0)
	// OutPut:
	// 	1 <nil>
	//result true
}
func BenchmarkPrepare1(b *testing.B) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lxq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var users []User
		var user User
		tx := db.Session(&gorm.Session{PrepareStmt: true})
		tx.First(&user, 1)
		tx.Where(fmt.Sprintf("name='junmo-%d' AND age>?", rand.Int()), 10).Find(&users)
	}

}

func BenchmarkPrepare2(b *testing.B) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lxq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var users []User
		var user User
		db.First(&user, 1)
		db.Where(fmt.Sprintf("name='junmo-%d' AND age>?", rand.Int()), 10).Find(&users)
	}
}
func BenchmarkPrepare3(b *testing.B) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lxq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		var users []User
		var user User
		tx := db.Session(&gorm.Session{PrepareStmt: true})
		tx.First(&user, 1)
		tx.Where("name=? AND age>?", fmt.Sprintf("junmo-%d", rand.Int()), 10).Find(&users)
	}
}

// 	BenchmarkSelectRedis
//	BenchmarkSelectRedis-8              7528            146485 ns/op
//	BenchmarkSelectLocal
//	BenchmarkSelectLocal-8            417232              2808 ns/op
func BenchmarkSelectRedis(b *testing.B) {
	b.StopTimer()
	RedisCacheInit("127.0.0.1", "6379", "")
	c := NewRedisCache()
	RegisterCache(c)
	db := NewDb()
	var user User
	stmt := db.DryRun().Find(&user, 1).Statement
	tag := "users000"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	}
}
func BenchmarkSelectLocal(b *testing.B) {
	b.StopTimer()
	c := NewLocalCache()
	RegisterCache(c)
	db := NewDb()
	var user User
	stmt := db.DryRun().Find(&user, 1).Statement
	tag := "users000"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		db.SetTag(tag).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&user)
	}
}
func BenchmarkHashMd5(b *testing.B) {
	str := "junmocsqjunmocsqjunmocsqjunmocsqjunmocsqjunmocsqjunmocsq"
	b.Run("hash-1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			md5Hash(str)
		}
	})
	b.Run("hash-2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			md5Hash2(str)
		}
	})
	b.Run("hash-3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			md5Hash3(str)
		}
	})
}
