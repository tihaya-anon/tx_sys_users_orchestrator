package cmd

import (
	"context"
	"os/signal"
	"syscall"
	"time"
	"users_orchestrator/config"
	"users_orchestrator/global/log"
	test_router "users_orchestrator/router/test"
	test_controller_builder "users_orchestrator/section/test/controller/builder"
	test_mapper_builder "users_orchestrator/section/test/mapper/builder"
	test_service_builder "users_orchestrator/section/test/service/builder"
	"users_orchestrator/server"

	"github.com/gin-gonic/gin"
)

func bindController() {
	testAMapper := test_mapper_builder.NewTestAMapperBuilder().Build()
	testAService := test_service_builder.NewTestAServiceBuilder().WithTestAMapper(testAMapper).Build()
	testAController := test_controller_builder.NewTestAControllerBuilder().WithTestAService(testAService).Build()
	testAController.Logger = log.GetLogger(24 * time.Hour)
	test_router.BindTestAController(testAController)
}

func startServer(publicPath, authPath string, engine *gin.Engine, timeOut time.Duration) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := server.NewServer()
	server.Setup(publicPath, authPath, engine)
	server.Run()

	<-ctx.Done()

	server.Stop(timeOut)
}

func Start() {
	logger := log.GetLogger(24 * time.Hour)
	logger.Infof("============= START =============")
	bindController()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	logger.Infof("activate profile: %s", config.Application.Env)
	logger.Infof("listen to: %s", config.Application.App.Uri)
	publicPath := "/api/v1/public"
	authPath := "/api/v1/auth"
	logger.Debugf("public path: %s", publicPath)
	logger.Debugf("auth path: %s", authPath)
	startServer(publicPath, authPath, engine, 5*time.Second)
}

func Stop() {
	log.GetLogger(24 * time.Hour).Infof("============= STOP =============")
}
