package user_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
  user_controller "users_orchestrator/section/user/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindUserController (ctrl *user_controller.UserController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/user"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/user"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}