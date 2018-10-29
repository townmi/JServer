package restful

import (
	"net/http"

	utils "../utils"
	"github.com/gin-gonic/gin"
)

// AccountLogin s
func AccountLogin(c *gin.Context) {
	account := map[string]string{
		"email":    "",
		"password": "",
	}
	isValidate := utils.ValidateRequestForm(c, account)
	if isValidate {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "ok",
		})
	}
}
