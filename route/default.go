package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("lxq")
		c.String(200, "Success lxq\n")
		c.Next()
		fmt.Println("csq")
		c.String(200, "Success csq\n")
	}
}
