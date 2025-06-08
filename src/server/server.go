package server

import (
	"context"
	"net/http"
	"time"
	"users_orchestrator/config"
	"users_orchestrator/global/log"
	"users_orchestrator/middleware"
	"users_orchestrator/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Setup(publicPath, authPath string, engine *gin.Engine) {
	publicRouterGroup := engine.Group(publicPath)
	authRouterGroup := engine.Group(authPath)
	authRouterGroup.Use(middleware.JwtMiddleware())

	for _, registerRouterFunc := range router.RegisterRouterFuncList {
		registerRouterFunc(publicRouterGroup, authRouterGroup)
	}

	s.server = &http.Server{
		Addr:    config.Application.App.Uri,
		Handler: engine,
	}
}

func (s *Server) Run() {
	logger := log.GetLogger(24 * time.Hour)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Start server failed: %s", err)
			return
		}
	}()
}

func (s *Server) Stop(timeOut time.Duration) {
	logger := log.GetLogger(24 * time.Hour)
	ctx, cancelShutdowm := context.WithTimeout(context.Background(), timeOut)
	defer cancelShutdowm()

	if err := s.server.Shutdown(ctx); err != nil {
		logger.Errorf("Server shutdown failed: %s", err)
		return
	}
	logger.Infof("Server shutdown successfully")
}
