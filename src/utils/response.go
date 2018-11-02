package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// StandardEmptyMessage s
func StandardEmptyMessage(key string) string {
	return GetLocalizer("zh", "zh").MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ResponseNotAllowedEmptyMessage",
			Other: "{{.Key}} is not allowed empty!",
		},
		TemplateData: map[string]string{
			"Key": key,
		},
	})
}

// StandardErrorMessage s
func StandardErrorMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " is error!",
		"code":    0,
	}
}

// StandardFailMessage s
func StandardFailMessage(err string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": err,
		"code":    0,
	}
}

// StandardSuccessMessage s
func StandardSuccessMessage(key string) interface{} {
	return gin.H{
		"status":  "success",
		"message": key + " success!",
		"code":    1,
	}
}
