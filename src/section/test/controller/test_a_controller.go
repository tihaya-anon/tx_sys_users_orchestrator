package test_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	test_service "users_orchestrator/section/test/service"
	"users_orchestrator/vo/resp"
)

type TestAController struct {
	TestAService test_service.TestAService
	Logger *logrus.Logger
}

func (ctrl *TestAController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `test-a`")
}