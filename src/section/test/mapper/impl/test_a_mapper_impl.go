package test_mapper_impl

import (
	test_mapper "users_orchestrator/section/test/mapper"

	"gorm.io/gorm"
)

type TestAMapperImpl struct{
	DB *gorm.DB
}

// INTERFACE
var _ test_mapper.TestAMapper = (*TestAMapperImpl)(nil)