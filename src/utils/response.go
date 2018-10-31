package utils

import "github.com/gin-gonic/gin"

// StandardEmptyMessage s
func StandardEmptyMessage(key string) string {
	return key + " is not allowed empty!"
}

// StandardErrorMessage s
func StandardErrorMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " is error!",
	}
}

// StandardFailMessage s
func StandardFailMessage(err string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": err,
	}
}

// StandardSuccessMessage s
func StandardSuccessMessage(key string) interface{} {
	return gin.H{
		"status":  "success",
		"message": key + " success!",
	}
}
