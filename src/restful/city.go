package restful

import (
	"net/http"

	models "../models"
	sqls "../sqls"
	"github.com/gin-gonic/gin"
)

// GetCityList s
func GetCityList(c *gin.Context) {
	ls := []models.City{}
	ls = sqls.QueryCityList()
	c.AsciiJSON(http.StatusOK, ls)
}
