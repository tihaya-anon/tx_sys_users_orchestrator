package audit_mapper_impl

import (
	audit_mapper "users_orchestrator/section/audit/mapper"

	"gorm.io/gorm"
)

type AuditMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ audit_mapper.AuditMapper = (*AuditMapperImpl)(nil)