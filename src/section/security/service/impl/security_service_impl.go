package security_service_impl

import (
	security_service "users_orchestrator/section/security/service"
	security_mapper "users_orchestrator/section/security/mapper"
)

type SecurityServiceImpl struct{
	SecurityMapper security_mapper.SecurityMapper
}

// INTERFACE
var _ security_service.SecurityService = (*SecurityServiceImpl)(nil)
