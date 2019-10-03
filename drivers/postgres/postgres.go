package postgres

import (
	"database/sql"
	"dimo-backend/config"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func ConnectAsDefault() *PostgresDB {
	var (
		host     = config.SqlHost
		port     = config.SqlPort
		username = config.SqlUsername
		password = config.SqlPassword
		dbname   = config.SqlDbname
	)
	return Connect(host, port, username, password, dbname)
}

func Connect(host string, port int, username, password, dbname string) *PostgresDB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	Postgres.SQL = db
	return Postgres
}
