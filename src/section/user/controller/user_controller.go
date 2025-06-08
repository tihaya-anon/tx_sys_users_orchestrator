package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	user_service "users_orchestrator/section/user/service"
	"users_orchestrator/vo/resp"
)

type UserController struct {
	UserService user_service.UserService
	Logger *logrus.Logger
}

func (ctrl *UserController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `user`")
}