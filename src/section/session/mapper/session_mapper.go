package session_mapper

//go:generate mockgen -source=session_mapper.go -destination=..\..\..\mock\session\mapper\session_mapper_mock.go -package=session_mapper_mock
type SessionMapper interface {
	// DEFINE METHODS
}