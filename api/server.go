package api

import (
	"github.com/gin-gonic/gin"
	db "gomarketplace/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error}
}
