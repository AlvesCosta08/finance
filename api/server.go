package api

import (
	db "github.com/AlvesCosta08/finance/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct{
	store *db.SQLStore
	router *gin.Engine
}


func NewServer(store *db.SQLStore)  *Server{
	server := &Server{store: store}	
	router := gin.Default()

	// Users
	router.POST("/user", server.createUser)                
	router.GET("/user/:username", server.getUser)        
	router.GET("/user/id/:id", server.getUserById)       
	router.GET("/user", server.getAllUsers)                
	router.PUT("/user/:id", server.updateUser)             
	router.DELETE("/user/:id", server.deleteUser)  


	//Categories
	router.POST("/cateory", server.createCategory)  
	router.GET("/category/id/:id", server.getCategory)  
	router.DELETE("/category/:id", server.deleteCategory)   

	server.router = router

	return server
}

func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"api has error:":err.Error(),
	}
}