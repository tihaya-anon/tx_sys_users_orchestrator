package audit_service

//go:generate mockgen -source=audit_service.go -destination=..\..\..\mock\audit\service\audit_service_mock.go -package=audit_service_mock
type AuditService interface {
	// DEFINE METHODS
}