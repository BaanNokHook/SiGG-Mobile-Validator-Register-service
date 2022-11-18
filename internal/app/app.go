package app

import (
	"fmt"
	"nextclan/validator-register/mobile-validator-register-service/config"
	v1 "nextclan/validator-register/mobile-validator-register-service/internal/controller/http/v1"
	usecase "nextclan/validator-register/mobile-validator-register-service/internal/usecase"
	"nextclan/validator-register/mobile-validator-register-service/internal/usecase/repo"
	"nextclan/validator-register/mobile-validator-register-service/pkg/httpserver"
	"nextclan/validator-register/mobile-validator-register-service/pkg/logger"
	mongodb "nextclan/validator-register/mobile-validator-register-service/pkg/mongo"
	"nextclan/validator-register/mobile-validator-register-service/pkg/pusher"
	"nextclan/validator-register/mobile-validator-register-service/pkg/redis"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	fmt.Println("Starting App...")
	deviceCache := redis.NewRedisClient(cfg.Redis.Addr, cfg.Password, cfg.Redis.RedisDeviceDB)
	mongoDb := mongodb.New(cfg.Mongo.ConnectionUri)
	pusherBeam := pusher.NewPusherClient(cfg.PusherBeam.InstanceId, cfg.PusherBeam.SecretKey)
	deviceRepository := repo.NewDeviceRepository(mongoDb.Client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.DeviceCollectionName))
	// Use case
	createMobileValidatorDeviceUseCase := usecase.NewCreateMobileValidatorDevice(l, deviceRepository, deviceCache)
	updateMobileValidatorDeviceStatusUseCase := usecase.NewUpdateMobileValidatorDeviceStatus(l, deviceRepository)
	pusherAuthenticationUseCase := usecase.NewPusherAuthentication(pusherBeam)
	// HTTP Server
	httpServer := initializeHttp(l, createMobileValidatorDeviceUseCase, updateMobileValidatorDeviceStatusUseCase, pusherAuthenticationUseCase, cfg)
	// Shutdown
	shutdownApplicationHandler(l, httpServer, mongoDb, deviceCache)
}

func initializeHttp(l *logger.Logger, cmv *usecase.CreateMobileValidatorDeviceUsecase, umv *usecase.UpdateMobileValidatorDeviceStatusUseCase, pa *usecase.PusherAuthenticationUsecase, cfg *config.Config) *httpserver.Server {
	handler := gin.New()
	handler.Use(cors.Default())
	v1.NewRouter(handler, cmv, umv, pa, l)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	return httpServer
}

func shutdownApplicationHandler(l *logger.Logger, httpServer *httpserver.Server, mongoDB *mongodb.MongoDB, redisDevice redis.RedisInterface) {
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	}
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	mongoDB.Close()
	redisDevice.Close()
}
