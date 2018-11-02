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
	r.Use(utils.CorsMiddleware())
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

	auth := r.Group("/auth")
	auth.Use(utils.JWTAuth())
	{
		auth.GET("/captcha", restful.CreateMathCaptcha)
	}

	goods := r.Group("/goods")
	{
		goods.POST("/create", restful.CreateGoods)
	}

	r.Run(":3000")
}
