package test_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
  test_controller "users_orchestrator/section/test/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindTestAController (ctrl *test_controller.TestAController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/test-a"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/test-a"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}