// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/fazilnbr/project-workey/pkg/api"
	"github.com/fazilnbr/project-workey/pkg/api/handler"
	"github.com/fazilnbr/project-workey/pkg/api/middleware"
	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/db"
	"github.com/fazilnbr/project-workey/pkg/repository"
	"github.com/fazilnbr/project-workey/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	sqlDB := db.ConnectDB(cfg)
	adminRepository := repository.NewAdminRepo(sqlDB)
	mailConfig := config.NewMailConfig()
	adminUseCase := usecase.NewAdminService(adminRepository, mailConfig)
	workerRepository := repository.NewWorkerRepo(sqlDB)
	workerUseCase := usecase.NewWorkerService(workerRepository)
	userRepository := repository.NewUserRepo(sqlDB)
	userUseCase := usecase.NewUserService(userRepository)
	jwtUseCase := usecase.NewJWTUserService()
	authUseCase := usecase.NewAuthService(adminRepository, workerRepository, userRepository, mailConfig, cfg)
	authHandler := handler.NewUserHandler(adminUseCase, workerUseCase, userUseCase, jwtUseCase, authUseCase)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	middlewareMiddleware := middleware.NewUserMiddileware(jwtUseCase)
	serverHTTP := api.NewServerHTTP(authHandler, adminHandler, middlewareMiddleware)
	return serverHTTP, nil
}
