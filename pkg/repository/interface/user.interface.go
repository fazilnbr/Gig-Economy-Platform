package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type UserRepository interface {
	InsertUser(login domain.Login) (int, error)
	FindUser(email string) (domain.UserResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	AddProfile(Profile domain.Profile, id int) (int, error)
	EditProfile(Profile domain.Profile, id int) (int, error)
	ChangePassword(changepassword string, id int) (int, error)
	// verifyPassword(changepassword string, id int) (int, error)
}
