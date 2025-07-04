package security_controller_builder

import (
  security_service "users_orchestrator/section/security/service"
  security_controller "users_orchestrator/section/security/controller"
)

func (builder *SecurityControllerBuilder) Build() *security_controller.SecurityController {
  if builder.isStrict && builder.securityController.SecurityService == nil {
    panic("`SecurityService` is required")
  }
  return builder.securityController
}

func (builder *SecurityControllerBuilder) WithSecurityService(securityService security_service.SecurityService) *SecurityControllerBuilder {
  builder.securityController.SecurityService = securityService
  return builder
}

// BUILDER
type SecurityControllerBuilder struct {
  isStrict bool
  securityController *security_controller.SecurityController
}

func NewSecurityControllerBuilder() *SecurityControllerBuilder {
  return &SecurityControllerBuilder{
    isStrict: false,
    securityController: &security_controller.SecurityController{},
  }
}

func (builder *SecurityControllerBuilder) UseStrict() *SecurityControllerBuilder { 
  builder.isStrict = true
  return builder
}