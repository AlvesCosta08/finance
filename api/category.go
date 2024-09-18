package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlvesCosta08/finance/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) getAllCategories(ctx *gin.Context) {
	categories, err := server.store.GetCategories(ctx,db.GetCategoriesParams{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}


type deleteCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteCategory(ctx, req.ID)
	if err != nil {				
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))		
		return
	}

	ctx.JSON(http.StatusOK,true)
}


