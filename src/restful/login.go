package restful

import (
	"net/http"

	sqls "../sqls"
	utils "../utils"
	"github.com/gin-gonic/gin"
)

// AccountLogin s
func AccountLogin(c *gin.Context) {
	account := map[string]string{
		"email":    "",
		"password": "",
		"captcha":  "",
	}
	err := utils.ValidateRequestForm(c, account)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}

	if !verfiyCaptcha(account["captcha"]) {
		c.JSON(http.StatusOK, utils.StandardErrorMessage("captcha"))
		return
	}

	validateErr := sqls.ValidateUser(account)
	if validateErr != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(validateErr.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ok",
	})
}
