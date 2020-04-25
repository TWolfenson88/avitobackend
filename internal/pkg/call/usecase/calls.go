package usecase

import (
	"avitocalls/internal/pkg/call"
	"avitocalls/internal/pkg/call/repository"
	"avitocalls/internal/pkg/db"
	"avitocalls/internal/pkg/models"
)

type callUseCase struct {
	rep call.Repository
}

func GetUseCase() call.UseCase {
	return &callUseCase{
		rep: repository.NewSqlCallRepository(db.ConnectToDB()),
	}
}


func (c callUseCase) GetReceiverObject(call models.Call) (models.Call, int, error) {
	// toDo normal implement
	return call, 200, nil
}

