package main

import (
	"./restful"
	utils "./utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.CheckDataBase()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	login := r.Group("/login")
	{
		login.POST("/account", restful.Login)
		login.POST("/sms", restful.Login)
		login.POST("/wechat", restful.Login)
	}

	r.GET("/city", restful.GetCityList)

	r.Run(":3000")
}
