package auth_service_impl

import (
	auth_service "users_orchestrator/section/auth/service"
	auth_mapper "users_orchestrator/section/auth/mapper"
)

type AuthServiceImpl struct{
	AuthMapper auth_mapper.AuthMapper
}

// INTERFACE
var _ auth_service.AuthService = (*AuthServiceImpl)(nil)
