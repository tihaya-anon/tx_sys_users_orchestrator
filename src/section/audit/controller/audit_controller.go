package audit_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	audit_service "users_orchestrator/section/audit/service"
	"users_orchestrator/vo/resp"
)

type AuditController struct {
	AuditService audit_service.AuditService
	Logger *logrus.Logger
}

func (ctrl *AuditController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `audit`")
}