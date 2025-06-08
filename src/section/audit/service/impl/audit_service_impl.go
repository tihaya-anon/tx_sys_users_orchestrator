package audit_service_impl

import (
	audit_service "users_orchestrator/section/audit/service"
	audit_mapper "users_orchestrator/section/audit/mapper"
)

type AuditServiceImpl struct{
	AuditMapper audit_mapper.AuditMapper
}

// INTERFACE
var _ audit_service.AuditService = (*AuditServiceImpl)(nil)
