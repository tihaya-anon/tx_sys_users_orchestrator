package user_controller_builder

import (
  user_service "users_orchestrator/section/user/service"
  user_controller "users_orchestrator/section/user/controller"
)

func (builder *UserControllerBuilder) Build() *user_controller.UserController {
  if builder.isStrict && builder.userController.UserService == nil {
    panic("`UserService` is required")
  }
  return builder.userController
}

func (builder *UserControllerBuilder) WithUserService(userService user_service.UserService) *UserControllerBuilder {
  builder.userController.UserService = userService
  return builder
}

// BUILDER
type UserControllerBuilder struct {
  isStrict bool
  userController *user_controller.UserController
}

func NewUserControllerBuilder() *UserControllerBuilder {
  return &UserControllerBuilder{
    isStrict: false,
    userController: &user_controller.UserController{},
  }
}

func (builder *UserControllerBuilder) UseStrict() *UserControllerBuilder { 
  builder.isStrict = true
  return builder
}