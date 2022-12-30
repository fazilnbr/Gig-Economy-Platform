package interfaces

import "github.com/fazilnbr/project-workey/pkg/domain"

type UserUseCase interface {
	CreateUser(newUser domain.Login) error
	FindUser(email string) (*domain.UserResponse, error)
	VerifyUser(email string, password string) error
}
