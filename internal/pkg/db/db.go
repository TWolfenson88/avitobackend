package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"sync"
)

var (
	//host     = "rc1b-9fc2waa12uuwznii.mdb.yandexcloud.net" // "localhost"
	//port     = 6432 // uint16(5432)
	//user     = "bakaef" // = os.Getenv("DB_USER")
	//password = "qwertyqwerty" // = os.Getenv("DB_PASSWORD")
	//dbname   = "calls" // = os.Getenv("DB_NAME")
	host     = "95.163.180.8" // "localhost"
	port     = 2000 // uint16(5432)
	user     = "postgres" // = os.Getenv("DB_USER")
	password = "qwerty" // = os.Getenv("DB_PASSWORD")
	dbname   = "postgres" // = os.Getenv("DB_NAME")
)


var db *pgx.ConnPool = nil
var syncOnce = sync.Once{}

func ConnectToDB() *pgx.ConnPool {
	syncOnce.Do(func() {
		pgxConfig := pgx.ConnConfig{
			Host:     host,
			Port:     uint16(port),
			Database: dbname,
			User:     user,
			Password: password,
		}
		pgxConnPoolConfig := pgx.ConnPoolConfig{
			ConnConfig: pgxConfig,
		}
		dbase, err := pgx.NewConnPool(pgxConnPoolConfig)
		if err != nil {
			log.Fatal("Connection to database was failed")
		}
		fmt.Println("connected")
		db = dbase
	})
	return db
}
