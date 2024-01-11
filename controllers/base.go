package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (bc BaseController) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Pong!")
}
