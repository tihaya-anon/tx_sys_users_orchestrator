package controller_uitl

import (
	"users_orchestrator/vo/req"
	"users_orchestrator/vo/resp/common"

	"github.com/gin-gonic/gin"
)

func BindPageReq(ctx *gin.Context) (*req.TPageReq, *common.ValidationError) {
	return BindValidation[req.TPageReq](ctx)
}
