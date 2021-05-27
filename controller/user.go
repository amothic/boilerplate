package controller

import (
	"net/http"

	"github.com/amothic/boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase *usecase.UserInteractor
}

func NewUserController(u *usecase.UserInteractor) *UserController {
	return &UserController{
		usecase: u,
	}
}

func (controller *UserController) GetAll(c *gin.Context) {
	users, err := controller.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	user, err := controller.usecase.CreateUser(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, user)
}
