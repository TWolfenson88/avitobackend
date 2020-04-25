package repository

import (
	"avitocalls/internal/pkg/call"
	"github.com/jackc/pgx"
)

type sqlCallRepository struct {
	db *pgx.ConnPool
}

func NewSqlCallRepository(db *pgx.ConnPool) call.Repository {
	return &sqlCallRepository{db: db}
}

//func (er *sqlCallRepository) DoPing() error {
//	err = er.db.Ping()
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("Ping db Success!")
//	}
//	return nil
//}

