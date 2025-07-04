package user_mapper_impl

import (
	user_mapper "users_orchestrator/section/user/mapper"

	"gorm.io/gorm"
)

type UserMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ user_mapper.UserMapper = (*UserMapperImpl)(nil)