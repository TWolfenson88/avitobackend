package user

import (
	// "avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
)

type Repository interface {
	GetAllUsers() ([]models.User, error)
	UserRegistration(user models.User) (int, int, error)
	// Authorize(form forms.LoginForm) (int, int, error)
	// AddEntering(session models.Session) (int, error)
	// GetUserByUID(user *models.User) (error)
}
