package driver

import (
	"github.com/amothic/boilerplate/adapter/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userController *controller.UserController
}

func NewServer(uc *controller.UserController) *Server {
	return &Server{
		userController: uc,
	}
}

func (server *Server) Run() {
	router := gin.Default()
	router.GET("/users", server.userController.GetAll)
	router.POST("/user", server.userController.CreateUser)
	router.Run()
}
