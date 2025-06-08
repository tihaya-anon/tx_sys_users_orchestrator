package audit_mapper

//go:generate mockgen -source=audit_mapper.go -destination=..\..\..\mock\audit\mapper\audit_mapper_mock.go -package=audit_mapper_mock
type AuditMapper interface {
	// DEFINE METHODS
}