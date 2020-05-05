package repository

import (
	"avitocalls/internal/pkg/call"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
	"github.com/jackc/pgx"
	"time"
)

type sqlCallRepository struct {
	db *pgx.ConnPool
}

func NewSqlCallRepository(db *pgx.ConnPool) call.Repository {
	return &sqlCallRepository{db: db}
}

func (er *sqlCallRepository) SaveCallStartingInfo(call forms.CallStartForm) (int, error) {
	var callid int
	sqlStatement := `INSERT INTO call (caller, answerer, start_time) 
	VALUES ( $1, $2, $3) 
	returning id;`
	err := er.db.QueryRow(sqlStatement,
		call.Caller,
		call.Answerer,
		call.TimeStart).
		Scan(&callid)
	if err != nil {
		return -1, err
	}
	return callid, nil
}

func (er *sqlCallRepository) SaveCallEndingInfo(call forms.CallEndForm) error {
	var err error
	sqlStatement := `UPDATE call SET end_time=$1, result=true WHERE id=$2`
	_, err = er.db.Exec(sqlStatement, call.TimeEnd, call.CallID)
	return err
}

func (er *sqlCallRepository) GetCallHistoryByID(userID int) ([]models.Call, error) {
	sqlStatement := `Select 
	c.id, c.caller, p1.name, c.answerer, p2.name, c.start_time, c.end_time, c.result 
	from call as c
	join profile p1
	on c.caller=p1.uid
	join profile p2
	on c.answerer=p2.uid
	where c.caller=$1 or c.answerer=$1`
	rows, err := er.db.Query(sqlStatement, userID)
	var calls []models.Call
	var endtime time.Time
	for rows.Next() {
		modelInfo := models.Call{}

		_ = rows.Scan(
			&modelInfo.CallID,
			&modelInfo.CallerID,
			&modelInfo.CallerName,
			&modelInfo.AnswererID,
			&modelInfo.AnswererName,
			&modelInfo.StartTime,
			&endtime,
			&modelInfo.Result,
			)
		if endtime.Sub(modelInfo.StartTime) > 0 {
			modelInfo.Duration = int(endtime.Sub(modelInfo.StartTime).Seconds())
		}
		calls = append(calls, modelInfo)
	}
	return calls, err
}

