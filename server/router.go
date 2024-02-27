package server

import (
	"github.com/gin-gonic/gin"
	"gomarketplace/controllers"
	"gomarketplace/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	ping := new(controllers.BaseController)

	router.GET("/ping", ping.Hello)
	router.Use(middlewares.AuthMiddleware())

	return router
}
