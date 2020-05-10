package user

import (
	"avitocalls/internal/pkg/models"
	// "github.com/gofrs/uuid"
)

type UseCase interface {
	InitUsers(users []models.User) ([]models.User, int, error)
	RegUser(user models.User) (int, int, error)
	ValidateLogin(user models.User) (int, int, error)

	// for Socket
	SetOnline(username string) error
	SetOffline(username string) error
}
