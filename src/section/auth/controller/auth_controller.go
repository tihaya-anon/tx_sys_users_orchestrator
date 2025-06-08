package auth_controller

import (
	auth_service "users_orchestrator/section/auth/service"
	"users_orchestrator/section/auth/vo"
	controller_uitl "users_orchestrator/util/controller"
	"users_orchestrator/vo/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	AuthService auth_service.AuthService
	Logger      *logrus.Logger
}

func (ctrl *AuthController) Login(ctx *gin.Context) *resp.TResponse {
	loginRequest, err := controller_uitl.BindValidation[vo.LoginRequest](ctx)
	if err != nil {
		return resp.NewResponse().ValidationError(err)
	}
	ctrl.Logger.Debug(loginRequest)
	return resp.NewResponse().SuccessWithData("Login")
}

func (ctrl *AuthController) Logout(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `auth`")
}
