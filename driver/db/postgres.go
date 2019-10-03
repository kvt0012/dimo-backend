package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func ConnectDefault() *PostgresDB {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	var (
		host, _     = os.LookupEnv("SQL_HOST")
		port, _     = os.LookupEnv("SQL_PORT")
		username, _ = os.LookupEnv("SQL_USERNAME")
		password, _ = os.LookupEnv("SQL_PASSWORD")
		dbname, _   = os.LookupEnv("SQL_DBNAME")
	)
	iport, _ := strconv.Atoi(port)
	return Connect(host, iport, username, password, dbname)
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
