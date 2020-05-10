package user

import (
	// "avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
)

type Repository interface {
	GetAllUsers() ([]models.User, error)
	UserRegistration(user models.User) (int, int, error)
	UserLogin(user models.User) (int, int, error)

	// for sock usage
	UpdateStatus(username string, status bool) error
}
