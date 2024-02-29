package api

import (
	"database/sql"
	"errors"
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
	var body createUserRequest
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

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getOneUser(ctx *gin.Context) {
	var uri getUserRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, uri.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getListUser(ctx *gin.Context) {
	var params listUserRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListUsersParams{
		Limit:  params.PageSize,
		Offset: (params.PageID - 1) * params.PageSize,
	}

	users, err := server.store.ListUsers(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
