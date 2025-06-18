package lib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const DB_URL = "IE2_DB_URL"
const DB_USR = "IE2_DB_USERNAME"
const DB_PWD = "IE2_DB_PASSWORD"
const DB_PORT = "IE2_DB_PORT"
const DB_DEF_PORT = "5432"

func CloseConn(conn *pgx.Conn) {

	log.Print("closing DB connection")

	conn.Close(context.Background())
}

func GetPostgresConn(dbname string) (*pgx.Conn, error) {

	dburl := os.Getenv(DB_URL)
	user := os.Getenv(DB_USR)
	pwd := os.Getenv(DB_PWD)
	port := os.Getenv(DB_PORT)

	if len(dbname) <= 0 {
		return nil, errors.New("db name is empty")
	}

	if len(dburl) <= 0 {
		return nil, errors.New("db url is empty")
	}

	if len(user) <= 0 {
		return nil, errors.New("db username is empty")
	}

	if len(pwd) <= 0 {
		return nil, errors.New("db pwd is empty")
	}

	if len(port) <= 0 {
		log.Print("DB Port number is empty. Using default 5432.")
		port = DB_DEF_PORT
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pwd, dburl, port, dbname)

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
