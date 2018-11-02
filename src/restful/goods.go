package restful

import (
	"net/http"

	sqls "../sqls"
	utils "../utils"
	"github.com/gin-gonic/gin"
)

// CreateGoods s
func CreateGoods(c *gin.Context) {
	goods := map[string]string{
		"name":    "",
		"sku":     "",
		"desc":    "",
		"picture": "",
		"price":   "",
	}

	err := utils.ValidateRequestForm(c, goods)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	createErr := sqls.CreateGoods(goods)
	if createErr != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(createErr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   goods,
		"code":   1,
	})
}
