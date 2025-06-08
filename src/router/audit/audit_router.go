package audit_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
  audit_controller "users_orchestrator/section/audit/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindAuditController (ctrl *audit_controller.AuditController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/audit"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/audit"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}