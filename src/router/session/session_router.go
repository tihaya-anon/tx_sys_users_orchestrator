package session_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
  session_controller "users_orchestrator/section/session/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindSessionController (ctrl *session_controller.SessionController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/session"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/session"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}