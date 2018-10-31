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
	err := utils.ValidateRequestForm(c, account)
	if err != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}
	if !verfiyCaptcha(account["captcha"]) {
		c.JSON(http.StatusOK, utils.StandardErrorMessage("captcha"))
		return
	}

	registerErr := sqls.RegisterUser(account)
	if registerErr != nil {
		c.JSON(http.StatusOK, utils.StandardFailMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.StandardSuccessMessage("register user"))

}
