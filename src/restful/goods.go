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
		"type":    "",
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

// GetGoodsList s
func GetGoodsList(c *gin.Context) {
	filter := map[string]string{
		"limit":  "",
		"offset": "",
		"sort":   "",
	}

	err := utils.ValidateRequestQuery(c, filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	rows, err := sqls.QueryGoodsList(filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rows,
		"code":   1,
	})
}

// CreateBrand s
func CreateBrand(c *gin.Context) {
	brand := map[string]string{
		"name": "",
		"logo": "",
	}

	err := utils.ValidateRequestForm(c, brand)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	createErr := sqls.CreateBrand(brand)
	if createErr != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(createErr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   brand,
		"code":   1,
	})
}

// GetBrandList s
func GetBrandList(c *gin.Context) {
	filter := map[string]string{
		"limit":  "",
		"offset": "",
		"sort":   "",
	}

	err := utils.ValidateRequestQuery(c, filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	rows, err := sqls.QueryBrandList(filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rows,
		"code":   1,
	})
}

// CreateGoodsType s
func CreateGoodsType(c *gin.Context) {
	goodsType := map[string]string{
		"name":     "",
		"picture":  "",
		"parentId": "",
	}

	err := utils.ValidateRequestForm(c, goodsType)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	createErr := sqls.CreateGoodsType(goodsType)
	if createErr != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(createErr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   goodsType,
		"code":   1,
	})
}

// GetGoodsTypeList s
func GetGoodsTypeList(c *gin.Context) {
	filter := map[string]string{
		"limit":  "",
		"offset": "",
		"sort":   "",
	}

	err := utils.ValidateRequestQuery(c, filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	rows, err := sqls.QueryGoodsTypeList(filter)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rows,
		"code":   1,
	})
}
