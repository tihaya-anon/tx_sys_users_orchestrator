package session_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	session_service "users_orchestrator/section/session/service"
	"users_orchestrator/vo/resp"
)

type SessionController struct {
	SessionService session_service.SessionService
	Logger *logrus.Logger
}

func (ctrl *SessionController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `session`")
}