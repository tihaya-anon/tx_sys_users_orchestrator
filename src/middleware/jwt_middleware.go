package middleware

import (
	"strings"
	"users_orchestrator/global/enum"
	"users_orchestrator/security"
	"users_orchestrator/vo/resp"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := resp.NewResponse()
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.MISSING_TOKEN, enum.MSG.MISSING_TOKEN, nil))
			return
		}
		isLegal, token := extractToken(token)
		if !isLegal || !security.CheckJWT(token) {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.INVALID_TOKEN, enum.MSG.INVALID_TOKEN, nil))
			return
		}
		ctx.Next()
	}
}

func extractToken(token string) (bool, string) {
	isLegal := strings.HasPrefix(token, "Bearer ") && len(strings.Split(token, " ")) == 2
	if isLegal {
		return true, strings.Split(token, " ")[1]
	}
	return false, ""
}
