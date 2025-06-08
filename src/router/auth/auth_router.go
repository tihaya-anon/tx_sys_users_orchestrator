package auth_router

import (
	"users_orchestrator/middleware"
	"users_orchestrator/router"
	auth_controller "users_orchestrator/section/auth/controller"
	"users_orchestrator/util"

	"github.com/gin-gonic/gin"
)

func BindAuthController(ctrl *auth_controller.AuthController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/auth"))
		authGroup := util.RoutesWrapper(authRouterGroup.Group("/auth"))
		authGroup.Use(middleware.JwtMiddleware())

		publicGroup.POST("/login", ctrl.Login)
		authGroup.POST("/logout", ctrl.Logout)
	})
}
