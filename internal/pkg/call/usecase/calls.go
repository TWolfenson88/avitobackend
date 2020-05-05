package usecase

import (
	"avitocalls/internal/pkg/call"
	"avitocalls/internal/pkg/call/repository"
	"avitocalls/internal/pkg/db"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
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

func (c callUseCase) SaveCallEnding(form forms.CallEndForm) (int, error) {
	var err error
	if form.Result{
		loc, _ := time.LoadLocation("Europe/Moscow")
		form.TimeEnd = time.Now().In(loc)
		err = c.rep.SaveCallEndingInfo(form)
	}
	return 200, err
}

func (c callUseCase) GetUserCallHistory(uid forms.CallHistReqForm) ([]models.Call, error) {
	var err error
	var calls []models.Call
	calls, err = c.rep.GetCallHistoryByID(uid.UID)
	return calls, err
}



