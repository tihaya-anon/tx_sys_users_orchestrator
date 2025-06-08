package session_service_builder

import (
	session_service "users_orchestrator/section/session/service"
	session_service_impl "users_orchestrator/section/session/service/impl"
	session_mapper "users_orchestrator/section/session/mapper"
)

func (builder *SessionServiceBuilder) Build() session_service.SessionService {
	if builder.isStrict && builder.sessionServiceImpl.SessionMapper == nil {
		panic("`SessionMapper` is required")
	}
	return builder.sessionServiceImpl
}

func (builder *SessionServiceBuilder) WithSessionMapper(mapper session_mapper.SessionMapper) *SessionServiceBuilder {
	builder.sessionServiceImpl.SessionMapper = mapper
	return builder
}

// BUILDER
type SessionServiceBuilder struct {
  isStrict bool
	sessionServiceImpl *session_service_impl.SessionServiceImpl
}

func NewSessionServiceBuilder() *SessionServiceBuilder {
	return &SessionServiceBuilder{
		sessionServiceImpl: &session_service_impl.SessionServiceImpl{},
	}
}

func (builder *SessionServiceBuilder) UseStrict() *SessionServiceBuilder { 
  builder.isStrict = true
  return builder
}