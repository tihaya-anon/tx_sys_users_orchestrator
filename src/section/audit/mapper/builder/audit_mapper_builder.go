package audit_mapper_builder

import (
	audit_mapper "users_orchestrator/section/audit/mapper"
	audit_mapper_impl "users_orchestrator/section/audit/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *AuditMapperBuilder) Build() audit_mapper.AuditMapper {
	return builder.auditMapperImpl
}

func (builder *AuditMapperBuilder) WithDB(DB *gorm.DB) *AuditMapperBuilder {
  builder.auditMapperImpl.DB = DB
  return builder
}

// BUILDER
type AuditMapperBuilder struct {
  isStrict bool
	auditMapperImpl *audit_mapper_impl.AuditMapperImpl
}

func NewAuditMapperBuilder() *AuditMapperBuilder {
	return &AuditMapperBuilder{
		auditMapperImpl: &audit_mapper_impl.AuditMapperImpl{},
	}
}

func (builder *AuditMapperBuilder) UseStrict() *AuditMapperBuilder { 
  builder.isStrict = true
  return builder
}