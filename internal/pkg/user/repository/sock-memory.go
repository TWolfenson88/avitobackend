package repository

func (er *sqlUserRepository) UpdateStatus(username string, status bool) error {
	sqlStatement := `UPDATE profile SET status=$1 WHERE name=$2;`
	_, err := er.db.Exec(sqlStatement, status, username)
	return err
}
