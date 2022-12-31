//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/project-workey/pkg/api"
	handler "github.com/fazilnbr/project-workey/pkg/api/middleware"
	config "github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/db"
	repository "github.com/fazilnbr/project-workey/pkg/repository"
	usecase "github.com/fazilnbr/project-workey/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	// fmt.Printf("\n\n\nv\n\n\n")
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepo,
		config.NewMailConfig,
		usecase.NewJWTUserService,
		usecase.NewUserService,
		usecase.NewAuthService,
		handler.NewUserHandler,
		http.NewServerHTTP)

	// fmt.Printf("\n\n\nbuild : %v\n\n\n", s)
	return &http.ServerHTTP{}, nil

}
