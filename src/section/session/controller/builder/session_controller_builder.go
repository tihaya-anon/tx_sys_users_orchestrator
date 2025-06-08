package session_controller_builder

import (
  session_service "users_orchestrator/section/session/service"
  session_controller "users_orchestrator/section/session/controller"
)

func (builder *SessionControllerBuilder) Build() *session_controller.SessionController {
  if builder.isStrict && builder.sessionController.SessionService == nil {
    panic("`SessionService` is required")
  }
  return builder.sessionController
}

func (builder *SessionControllerBuilder) WithSessionService(sessionService session_service.SessionService) *SessionControllerBuilder {
  builder.sessionController.SessionService = sessionService
  return builder
}

// BUILDER
type SessionControllerBuilder struct {
  isStrict bool
  sessionController *session_controller.SessionController
}

func NewSessionControllerBuilder() *SessionControllerBuilder {
  return &SessionControllerBuilder{
    isStrict: false,
    sessionController: &session_controller.SessionController{},
  }
}

func (builder *SessionControllerBuilder) UseStrict() *SessionControllerBuilder { 
  builder.isStrict = true
  return builder
}