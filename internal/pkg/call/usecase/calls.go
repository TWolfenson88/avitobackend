package usecase

import (
	"avitocalls/internal/pkg/call"
	"avitocalls/internal/pkg/call/repository"
	"avitocalls/internal/pkg/db"
	"avitocalls/internal/pkg/forms"
	"time"
)



type callUseCase struct {
	rep call.Repository
}

func GetUseCase() call.UseCase {
	return &callUseCase{
		rep: repository.NewSqlCallRepository(db.ConnectToDB()),
	}
}

func (c callUseCase) SaveCallStarting(call forms.CallStartForm) (int, error) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	call.TimeStart = time.Now().In(loc)
	callid, err := c.rep.SaveCallStartingInfo(call)
	return callid, err
}

func (c callUseCase) SaveCallEnding(call forms.CallEndForm) (int, error) {
	var err error
	if call.Result{
		loc, _ := time.LoadLocation("Europe/Moscow")
		call.TimeEnd = time.Now().In(loc)
		err = c.rep.SaveCallEndingInfo(call)
	}
	return 200, err
}



