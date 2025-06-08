package session_service_impl

import (
	session_service "users_orchestrator/section/session/service"
	session_mapper "users_orchestrator/section/session/mapper"
)

type SessionServiceImpl struct{
	SessionMapper session_mapper.SessionMapper
}

// INTERFACE
var _ session_service.SessionService = (*SessionServiceImpl)(nil)
