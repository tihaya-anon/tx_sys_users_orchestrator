package audit_controller_builder

import (
  audit_service "users_orchestrator/section/audit/service"
  audit_controller "users_orchestrator/section/audit/controller"
)

func (builder *AuditControllerBuilder) Build() *audit_controller.AuditController {
  if builder.isStrict && builder.auditController.AuditService == nil {
    panic("`AuditService` is required")
  }
  return builder.auditController
}

func (builder *AuditControllerBuilder) WithAuditService(auditService audit_service.AuditService) *AuditControllerBuilder {
  builder.auditController.AuditService = auditService
  return builder
}

// BUILDER
type AuditControllerBuilder struct {
  isStrict bool
  auditController *audit_controller.AuditController
}

func NewAuditControllerBuilder() *AuditControllerBuilder {
  return &AuditControllerBuilder{
    isStrict: false,
    auditController: &audit_controller.AuditController{},
  }
}

func (builder *AuditControllerBuilder) UseStrict() *AuditControllerBuilder { 
  builder.isStrict = true
  return builder
}