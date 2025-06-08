package user_service_impl

import (
	user_service "users_orchestrator/section/user/service"
	user_mapper "users_orchestrator/section/user/mapper"
)

type UserServiceImpl struct{
	UserMapper user_mapper.UserMapper
}

// INTERFACE
var _ user_service.UserService = (*UserServiceImpl)(nil)
