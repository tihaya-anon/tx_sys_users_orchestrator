package test

import (
	"testing"
	"users_orchestrator/config"
	"users_orchestrator/util/gen"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Gen(t *testing.T) {
	gen.Generate("users_orchestrator", []string{"user"})
}
func Test_GenQuery(t *testing.T) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		t.Errorf("connect to database failed: %v", err)
	}
	gen.GenerateQuery([]string{"user"}, gormDB)
}

func Test_MVC(t *testing.T) {
	Test_GenController(t)
	Test_GenService(t)
	Test_GenMapper(t)
}
func Test_GenController(t *testing.T) {
	gen.GenerateGinController("users_orchestrator", "section", "test", []string{"test_a"})
	gen.GenerateGinController("users_orchestrator", "section", "auth", []string{"auth"})
	gen.GenerateGinController("users_orchestrator", "section", "user", []string{"user"})
	gen.GenerateGinController("users_orchestrator", "section", "security", []string{"security"})
	gen.GenerateGinController("users_orchestrator", "section", "session", []string{"session"})
	gen.GenerateGinController("users_orchestrator", "section", "audit", []string{"audit"})
}

func Test_GenService(t *testing.T) {
	gen.GenerateService("users_orchestrator", "section", "test", []string{"test_a"})
	gen.GenerateService("users_orchestrator", "section", "auth", []string{"auth"})
	gen.GenerateService("users_orchestrator", "section", "user", []string{"user"})
	gen.GenerateService("users_orchestrator", "section", "security", []string{"security"})
	gen.GenerateService("users_orchestrator", "section", "session", []string{"session"})
	gen.GenerateService("users_orchestrator", "section", "audit", []string{"audit"})
}

func Test_GenMapper(t *testing.T) {
	gen.GenerateMapper("users_orchestrator", "section", "test", []string{"test_a"})
	gen.GenerateMapper("users_orchestrator", "section", "auth", []string{"auth"})
	gen.GenerateMapper("users_orchestrator", "section", "user", []string{"user"})
	gen.GenerateMapper("users_orchestrator", "section", "security", []string{"security"})
	gen.GenerateMapper("users_orchestrator", "section", "session", []string{"session"})
	gen.GenerateMapper("users_orchestrator", "section", "audit", []string{"audit"})
}
