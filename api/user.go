package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/AlvesCosta08/finance/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}


func (server *Server) createUser(ctx *gin.Context)  {
	var req createUserRequest

    err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg:= db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
	}

	user, err := server.store.CreateUser(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK,user)
}

type getUserRequest struct{
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context)  {
	var req getUserRequest

    err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUser(ctx,req.Username)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return	
		}
		    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,user)
}

type getUserByIdRequest struct{
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getUserById(ctx *gin.Context)  {
	var req getUserByIdRequest

    err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUserByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return	
		}
		    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,user)
}

func (server *Server) getAllUsers(ctx *gin.Context) {
    users, err := server.store.ListUsers(ctx) 
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    ctx.JSON(http.StatusOK, users)
}

type UpdateUserRequest struct {
    ID       int32  `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

func (server *Server) updateUser(ctx *gin.Context) {
    var req UpdateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    // Converta UpdateUserRequest para db.UpdateUserParams
    dbParams := db.UpdateUserParams{
        ID:       int32(id),
        Username: req.Username,
        Password: req.Password,
        Email:    req.Email,
    }

    err = server.store.UpdateUser(ctx, dbParams)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User updated"})
}

type getUseDeleteRequest struct{
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req getUseDeleteRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, req.ID)
	if err != nil {				
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))		
		return
	}

	ctx.JSON(http.StatusOK,true)
}

