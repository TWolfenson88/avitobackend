package user

import (
	"avitocalls/internal/pkg/models"
	// "github.com/gofrs/uuid"
)

type UseCase interface {
	InitUsers(users []models.User) ([]models.User, int, error)
	RegUser(user models.User) (int, int, error)
	// LoginUser(loginForm forms.LoginForm) (int, int, error)
	// LogSession(session models.Session) (int, error)
	// FindUser(user *models.User) (int, error)
}

//type MemUseCase interface {
//	CreateSession(ip string, agent string, uid int) (models.Session, error)
//	ValidateSession(sessID uuid.UUID, agent string) (forms.CheckSessionAnswer, int, error)
//}

