package api

import "github.com/gin-gonic/gin"

type createUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	SurName   string `json:"sur_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (server *Server) createUser(ctx *gin.Context) {

}
