package session_mapper_impl

import (
	session_mapper "users_orchestrator/section/session/mapper"

	"gorm.io/gorm"
)

type SessionMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ session_mapper.SessionMapper = (*SessionMapperImpl)(nil)