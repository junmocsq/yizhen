package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Name      string    `json:"name" form:"name" binding:"required"`
	Passwd    string    `json:"passwd" form:"passwd" binding:"required,alphanum"`
	Age       int       `json:"age"`
	IsLogin   bool      `json:"is_login"`
	CreatedAt time.Time `json:"created_at"`
}

func V1User(r *gin.Engine) {
	un := r.Group("/user")
	{
		un.GET("/:id", func(c *gin.Context) {
			fmt.Println("query", c.Query("name"))
			fmt.Println("postform", c.PostForm("name"))
			var user User
			fmt.Println(c.Bind(&user))
			c.String(200, fmt.Sprintf("%#v", user))

			c.String(200, c.FullPath())
		})

		un.POST("/:id", func(c *gin.Context) {
			fmt.Println("query", c.Query("name"))
			fmt.Println("postform", c.PostForm("name"))

			var user User
			fmt.Println(c.Bind(&user))
			c.String(200, fmt.Sprintf("%#v", user))

			c.String(200, c.FullPath())
		})
		un.GET("/ids/:id", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.POST("/login", func(c *gin.Context) {

			c.String(200, c.FullPath())
		})
		un.POST("/reg", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.GET("/phoneCode/:phone", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.POST("/phoneCode/:phone", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.GET("/emailCode", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.POST("/emailCode", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.GET("/findPasswd", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		un.POST("/findPasswd", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
	}
	u := r.Group("/user", UserCheck())
	{
		u.GET("/detail/:id", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		u.PUT("/detail/:id", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		// 电话只能绑定和换绑
		u.PUT("/swapPhone", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
		u.PUT("/bind", func(c *gin.Context) {
			c.String(200, c.FullPath())
		})
	}
}
