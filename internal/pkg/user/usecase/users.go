package usecase

import (
	"avitocalls/internal/pkg/db"
	// "avitocalls/internal/pkg/forms"
	_ "avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
	"avitocalls/internal/pkg/user"
	"avitocalls/internal/pkg/user/repository"
	_ "fmt"
	"net/http"
)

type userUseCase struct {
	rep user.Repository
}

func GetUseCase() user.UseCase {
	return &userUseCase{
		rep: repository.NewSqlUserRepository(db.ConnectToDB()),
	}
}


func (uc *userUseCase) InitUsers(users []models.User) ([]models.User, int, error) {
	users, err := uc.rep.GetAllUsers()
	if err != nil {
		var bad []models.User
		return bad, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
}


func (uc *userUseCase) RegUser(user models.User) (int, int, error) {
	status, uid, err := uc.rep.UserRegistration(user)
	if status != http.StatusOK {
		return uid, status, err
	}
	return uid, http.StatusOK, nil
}

func (uc *userUseCase) ValidateLogin(user models.User) (int, int, error) {
	status, uid, err := uc.rep.UserLogin(user)
	if status != http.StatusOK {
		return uid, status, err
	}
	return uid, http.StatusOK, nil
}
