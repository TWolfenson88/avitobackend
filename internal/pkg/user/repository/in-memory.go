package repository
//
//import (
//	"counity/internal/pkg/models"
//	"counity/internal/pkg/user"
//	"fmt"
//	"github.com/gofrs/uuid"
//	"github.com/jackc/pgx"
//	"net/http"
//)
//
//type inMemoryUserRepository struct {
//	db *pgx.ConnPool
//}
//
//func NewInMemoryUserRepository(db *pgx.ConnPool) user.MemRepository {
//	return &inMemoryUserRepository{db: db}
//}
//
//func (mmr *inMemoryUserRepository) SelectSession(sessID uuid.UUID) (models.Session, error) {
//	var session models.Session
//	sqlStatement := `select sess_id, user_id, user_agent, add_time from session where sess_id = $1`
//	err := mmr.db.QueryRow(sqlStatement, sessID).Scan(
//		&session.SessID, &session.UserID, &session.UserAgent, &session.AddTime)
//	return session, err
//}
//
//func (mmr *inMemoryUserRepository) InitSession(session *models.Session) (int, error) {
//	sqlStatement := `insert into session (user_id, user_agent) values ($1, $2) returning sess_id, add_time`
//	err := mmr.db.QueryRow(sqlStatement, session.UserID, session.UserAgent).Scan(&session.SessID, &session.AddTime)
//	if err != nil {
//		fmt.Println("troubles with db")  // toDo log here too
//		return http.StatusInternalServerError, err
//	}
//
//	return http.StatusOK, nil
//}
