package repository

import "net/http"

func (er *sqlUserRepository) CheckConflicts(name string) (int, error) {
	sqlStatement := `SELECT count(*) cnt FROM profile WHERE name=$1;`
	cnt := 0
	row := er.db.QueryRow(sqlStatement, name)
	err := row.Scan(&cnt)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if cnt > 0 {
		return http.StatusConflict, nil
	}
	return http.StatusOK, nil
}
