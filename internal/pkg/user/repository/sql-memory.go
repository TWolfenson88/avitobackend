package repository

import (
	// "avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
	"avitocalls/internal/pkg/security"
	"avitocalls/internal/pkg/user"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	// "fmt"
	"github.com/jackc/pgx"
	// "golang.org/x/crypto/bcrypt"
	"net/http"
)

type sqlUserRepository struct {
	db *pgx.ConnPool
}

func NewSqlUserRepository(db *pgx.ConnPool) user.Repository {
	return &sqlUserRepository{db: db}
}


func (er *sqlUserRepository) GetAllUsers() ([]models.User, error) {
	sqlStatement := `SELECT uid, name, status FROM profile;`
	rows, err := er.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		modelInfo := models.User{}
		err = rows.Scan(
			&modelInfo.Uid,
			&modelInfo.Name,
			//&modelInfo.Email,
			//&modelInfo.Identify,
			&modelInfo.StatusOnline,
			)
		if err != nil {
			return nil, err
		}
		users = append(users, modelInfo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}


func (er *sqlUserRepository) UserRegistration(user models.User) (int, int, error) {
	check, err := er.CheckConflicts(user.Name)
	if err != nil {
		return check, -1, err
	}
	if check != 200 {
		return http.StatusConflict, -1, nil
	}
	hash, err := security.EncryptPassword(user.Password)
	// todo normal error catch (x3)
	if err != nil {
		return http.StatusInternalServerError, -1, err
	}
	//sqlStatement := `insert into session (user_id, user_agent) values ($1, $2) returning sess_id, add_time`
	//err := mmr.db.QueryRow(sqlStatement, session.UserID, session.UserAgent).Scan(&session.SessID, &session.AddTime)
	var uid int
	sqlStatement := `
	INSERT INTO profile (name, password) 
	VALUES ($1, $2)
	returning uid;
	`
	err = er.db.QueryRow(sqlStatement,
		user.Name,
		hash).Scan(&uid)
	fmt.Println(uid)
	// todo normal error catch
	if err != nil {
		return http.StatusInternalServerError, -1, err
	}
	return http.StatusOK, uid, nil
}

func (er *sqlUserRepository) UserLogin(user models.User) (int, int, error) {
	cnt := 0
	sqlStatement := `SELECT count(*) cnt FROM profile WHERE name=$1;`
	row := er.db.QueryRow(sqlStatement, user.Name)
	err := row.Scan(&cnt)
	if cnt == 0 {
		return http.StatusConflict, -1, nil
	}
	uid := -1
	var pass []byte
	sqlStatement = `SELECT uid, password FROM profile WHERE name=$1;`
	row = er.db.QueryRow(sqlStatement, user.Name)
	err = row.Scan(&uid, &pass)
	if err != nil {
		return http.StatusInternalServerError, -1, err
	}
	err = bcrypt.CompareHashAndPassword(pass, []byte(user.Password))
	if err == nil {
		return http.StatusOK, uid, nil
	} else {
		return http.StatusForbidden, -1, nil
	}
}
