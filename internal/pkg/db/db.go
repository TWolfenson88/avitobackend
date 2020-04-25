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
	// sslmode = "verify-full"
	// sslrootcert = "/home/username/.postgresql/root.crt"

	// please do not use this db :)
	host     = "95.163.180.8"
	port     = 2000
	user     = "postgres"
	password = "qwerty"
	dbname   = "postgres"
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
