package test_service_impl

import (
	test_service "users_orchestrator/section/test/service"
	test_mapper "users_orchestrator/section/test/mapper"
)

type TestAServiceImpl struct{
	TestAMapper test_mapper.TestAMapper
}

// INTERFACE
var _ test_service.TestAService = (*TestAServiceImpl)(nil)
