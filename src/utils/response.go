package utils

import "github.com/gin-gonic/gin"

func StandardEmptyMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " is not allowed empty!",
	}
}
