package security_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
  security_controller "users_orchestrator/section/security/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindSecurityController (ctrl *security_controller.SecurityController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/security"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/security"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}