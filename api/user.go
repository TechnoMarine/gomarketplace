package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "gomarketplace/db/sqlc"
	"log"
	"net/http"
	"time"
)

type createUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	SurName   string `json:"surName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	body := createUserRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		SurName:   body.SurName,
		Email:     body.Email,
		Password:  body.Password,
		Address:   body.Address,
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{},
	}

	row, err := server.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal("Create user error")
		return
	}

	ctx.JSON(http.StatusCreated, row)
}

func (server *Server) getOneUser(ctx *gin.Context) {

}
