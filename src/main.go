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
		login.POST("/account", restful.AccountLogin)
		login.POST("/sms", restful.Login)
		login.POST("/wechat", restful.Login)
	}

	register := r.Group("/register")
	{
		register.POST("account", restful.AccountRegister)
	}

	public := r.Group("/public")
	{
		public.GET("/city", restful.GetCityList)
		public.GET("/captcha", restful.CreateMathCaptcha)
	}

	r.Run(":3000")
}
