package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/middleware"
	"github.com/ryanma3003/hris/servers"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
	db.CasbinAdapter()
	db.CasbinEnforcer()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.POST("login", controllers.Login)
		v1.GET("casbin", controllers.GetFrontendPermission)
		v1.GET("validate", middleware.Authorize("resource", "*"), controllers.Validate)

		servers.InitializeRoutesEmployee(v1.Group("main"))
		servers.InitializeRoutesMpp(v1.Group("mpp"))
		servers.InitializeRoutesCompliance(v1.Group("compliance"))
		servers.InitializeRoutesService(v1.Group("service"))
		servers.InitializeRoutesPerformance(v1.Group("performance"))

		v1.POST("/logout", controllers.Logout)
	}

	router.Run()
}
