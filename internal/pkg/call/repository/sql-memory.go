package repository

import (
	"avitocalls/internal/pkg/call"
	"avitocalls/internal/pkg/forms"
	"github.com/jackc/pgx"
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

