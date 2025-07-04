package test_service_builder

import (
	test_service "users_orchestrator/section/test/service"
	test_service_impl "users_orchestrator/section/test/service/impl"
	test_mapper "users_orchestrator/section/test/mapper"
)

func (builder *TestAServiceBuilder) Build() test_service.TestAService {
	if builder.isStrict && builder.testAServiceImpl.TestAMapper == nil {
		panic("`TestAMapper` is required")
	}
	return builder.testAServiceImpl
}

func (builder *TestAServiceBuilder) WithTestAMapper(mapper test_mapper.TestAMapper) *TestAServiceBuilder {
	builder.testAServiceImpl.TestAMapper = mapper
	return builder
}

// BUILDER
type TestAServiceBuilder struct {
  isStrict bool
	testAServiceImpl *test_service_impl.TestAServiceImpl
}

func NewTestAServiceBuilder() *TestAServiceBuilder {
	return &TestAServiceBuilder{
		testAServiceImpl: &test_service_impl.TestAServiceImpl{},
	}
}

func (builder *TestAServiceBuilder) UseStrict() *TestAServiceBuilder { 
  builder.isStrict = true
  return builder
}