package server

import (
	"gomarketplace/controllers"
	"gomarketplace/middlewares"

	"github.com/gin-gonic/gin"
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
