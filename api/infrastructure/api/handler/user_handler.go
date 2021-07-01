package handler

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashita1023/clean-architecture/domain/model"
	"github.com/mashita1023/clean-architecture/interface/controllers"
)

type userHandler struct {
  userController controllers.UserController
}

type UserHandler interface {
  CreateUser(r gin.Context) error
  GetUsers(r gin.Context) error
}

func NewUserHandler(uc controllers.UserController) UserHandler {
  return &userHandler{userController: uc}
}

func (uh *userHandler) CreateUser(r gin.Context) error {
  req := &model.User{}

  // bind & validate
  if err := r.Bind(req); err != nil {
    return r.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
  }

  ctx := 

}
