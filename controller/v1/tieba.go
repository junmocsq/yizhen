package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/junmocsq/yizhen/models/tieba"
	"github.com/junmocsq/yizhen/utils"
	"net/http"
)

type tb struct {
}

func NewTB() *tb {
	return &tb{}
}
func (t *tb) Hello(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
}

// PingExample godoc
// @Summary 贴吧信息
// @Schemes
// @Description do ping
// @Tags Tieba
// @Accept json
// @Produce json
// @Param   name     query    string true "贴吧名称" default("赵小凡")
// @Success 200 {object} tieba.Tieba
// @Router /v1/tieba/tiebaByName [get]
func (t *tb) TiebaByName(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
}

// @Summary 贴吧信息
// @Schemes
// @Description do ping
// @Tags Tieba
// @Accept json
// @Produce json
// @Param   id     query    int true "贴吧ID" default(1)
// @Success 200 {object} tieba.Tieba
// @Router /v1/tieba/tiebaById [get]
func (t *tb) TiebaById(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
}

// @Description 获取贴吧列表
// @tags Tieba
// @Accept  json
// @Produce  json
// @Param   page     query    int     true        "分页" default(1)
// @Param   size      query    int     false        "每页条数" default(10)
// @Success 200 {array} []tieba.Tieba	"ok"
// @Failure 400 {string} failed "failed!!"
// @Router /v1/tieba/tiebaList [get]
func (t *tb) TiebaList(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	tieba.NewTieba().List(1, 10)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  utils.Msg(200),
		"data": tieba.NewTieba().List(1, 10),
	})
	fmt.Printf("ids: %v; names: %v", ids, names)
}
