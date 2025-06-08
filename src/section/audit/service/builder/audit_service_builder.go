package audit_service_builder

import (
	audit_service "users_orchestrator/section/audit/service"
	audit_service_impl "users_orchestrator/section/audit/service/impl"
	audit_mapper "users_orchestrator/section/audit/mapper"
)

func (builder *AuditServiceBuilder) Build() audit_service.AuditService {
	if builder.isStrict && builder.auditServiceImpl.AuditMapper == nil {
		panic("`AuditMapper` is required")
	}
	return builder.auditServiceImpl
}

func (builder *AuditServiceBuilder) WithAuditMapper(mapper audit_mapper.AuditMapper) *AuditServiceBuilder {
	builder.auditServiceImpl.AuditMapper = mapper
	return builder
}

// BUILDER
type AuditServiceBuilder struct {
  isStrict bool
	auditServiceImpl *audit_service_impl.AuditServiceImpl
}

func NewAuditServiceBuilder() *AuditServiceBuilder {
	return &AuditServiceBuilder{
		auditServiceImpl: &audit_service_impl.AuditServiceImpl{},
	}
}

func (builder *AuditServiceBuilder) UseStrict() *AuditServiceBuilder { 
  builder.isStrict = true
  return builder
}