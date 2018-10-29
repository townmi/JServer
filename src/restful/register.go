package restful

import (
	"net/http"

	sqls "../sqls"
	utils "../utils"
	"github.com/gin-gonic/gin"
)

// AccountRegisterForm s
type AccountRegisterForm struct {
	Email    string
	Password string
}

// AccountRegister s
func AccountRegister(c *gin.Context) {
	account := map[string]string{
		"email":    "",
		"password": "",
		"captcha":  "",
	}
	isValidate := utils.ValidateRequestForm(c, account)

	if isValidate {
		if !verfiyCaptcha(account["captcha"]) {
			c.JSON(http.StatusOK, utils.StandardErrorMessage("captcha"))
			return
		}
		success, err := sqls.RegisterUser(account)
		if err != nil || !success {
			c.JSON(http.StatusOK, utils.StandardFailMessage("register user"))
			return
		}
		c.JSON(http.StatusOK, utils.StandardSuccessMessage("register user"))
	}
}
