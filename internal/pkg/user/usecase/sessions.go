package usecase
//
//import (
//	"counity/internal/pkg/db"
//	"counity/internal/pkg/forms"
//	"counity/internal/pkg/models"
//	"counity/internal/pkg/user"
//	"counity/internal/pkg/user/repository"
//	"github.com/gofrs/uuid"
//	"net/http"
//)
//
//type sessionUseCase struct {
//	rep user.MemRepository
//}
//
//func GetSessUseCase() user.MemUseCase {
//	return &sessionUseCase{
//		rep: repository.NewInMemoryUserRepository(db.ConnectToDB()),
//	}
//}
//
//
//
//func (muc *sessionUseCase) ValidateSession(sessID uuid.UUID, agent string) (forms.CheckSessionAnswer, int, error) {
//	session, err := muc.rep.SelectSession(sessID)
//	if err != nil {
//		// toDo log this db shit!
//		return forms.DbTrouble, -1, err
//	}
//	if session.UserAgent != agent {
//		return forms.BadUserAgent, -1, nil
//	}
//	//if time.Now().Sub(session.AddTime).Hours() > 24 {
//	//	return forms.Expired, -1, nil
//	//}
//	return forms.OK, session.UserID, nil
//}
//
//func (muc *sessionUseCase) CreateSession(ip string, agent string, uid int) (models.Session, error) {
//	newSession := models.Session{
//		UserID:    uid,
//		UserAgent: agent,
//		IpAddress: ip,
//	}
//	_, err := muc.rep.InitSession(&newSession)
//	//if err != nil {
//	//	return newSession, err
//	//}
//	return newSession, err  // nil
//}
//
//func (uc *userUseCase) LogSession(session models.Session) (int, error) {
//	status, err := uc.rep.AddEntering(session)
//	if err != nil {
//		// toDo log error here
//		return status, err
//	}
//	return http.StatusOK, nil
//}
