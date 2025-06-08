package session_mapper_builder

import (
	session_mapper "users_orchestrator/section/session/mapper"
	session_mapper_impl "users_orchestrator/section/session/mapper/impl"
	
	"gorm.io/gorm"
)

func (builder *SessionMapperBuilder) Build() session_mapper.SessionMapper {
	return builder.sessionMapperImpl
}

func (builder *SessionMapperBuilder) WithDB(DB *gorm.DB) *SessionMapperBuilder {
  builder.sessionMapperImpl.DB = DB
  return builder
}

// BUILDER
type SessionMapperBuilder struct {
  isStrict bool
	sessionMapperImpl *session_mapper_impl.SessionMapperImpl
}

func NewSessionMapperBuilder() *SessionMapperBuilder {
	return &SessionMapperBuilder{
		sessionMapperImpl: &session_mapper_impl.SessionMapperImpl{},
	}
}

func (builder *SessionMapperBuilder) UseStrict() *SessionMapperBuilder { 
  builder.isStrict = true
  return builder
}