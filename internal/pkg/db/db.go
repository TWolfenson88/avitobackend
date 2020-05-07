package db


import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"sync"
)

// not secured? come on, man, welcome
var (
	//host     = "database-1.czbpipfhq5vi.eu-west-2.rds.amazonaws.com"
	//port     = 7000
	//user     = "postgres"
	//password = "qwertyqwerty"
	//dbname   = "postgres"
	host     = "84.201.143.114"
	port     = 3000
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

