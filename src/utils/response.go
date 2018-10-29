package utils

import "github.com/gin-gonic/gin"

// StandardEmptyMessage s
func StandardEmptyMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " is not allowed empty!",
	}
}

// StandardErrorMessage s
func StandardErrorMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " is error!",
	}
}

// StandardFailMessage s
func StandardFailMessage(key string) interface{} {
	return gin.H{
		"status":  "fail",
		"message": key + " fail!",
	}
}

// StandardSuccessMessage s
func StandardSuccessMessage(key string) interface{} {
	return gin.H{
		"status":  "success",
		"message": key + " success!",
	}
}
