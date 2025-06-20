package auth_controller_builder

import (
  auth_service "users_orchestrator/section/auth/service"
  auth_controller "users_orchestrator/section/auth/controller"
)

func (builder *AuthControllerBuilder) Build() *auth_controller.AuthController {
  if builder.isStrict && builder.authController.AuthService == nil {
    panic("`AuthService` is required")
  }
  return builder.authController
}

func (builder *AuthControllerBuilder) WithAuthService(authService auth_service.AuthService) *AuthControllerBuilder {
  builder.authController.AuthService = authService
  return builder
}

// BUILDER
type AuthControllerBuilder struct {
  isStrict bool
  authController *auth_controller.AuthController
}

func NewAuthControllerBuilder() *AuthControllerBuilder {
  return &AuthControllerBuilder{
    isStrict: false,
    authController: &auth_controller.AuthController{},
  }
}

func (builder *AuthControllerBuilder) UseStrict() *AuthControllerBuilder { 
  builder.isStrict = true
  return builder
}