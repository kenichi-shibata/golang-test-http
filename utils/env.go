package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// ENV (POSTGRES OR MYSQL)
// defaults to sqlite for local dev running as volume mounted
// psql -h host -U username
// example connString: "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

var database *sql.DB

const postgresCreateTable = `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, birthdate TEXT)`
const sqlite3CreateTable = `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate TEXT)`

func SQLOpen() (database *sql.DB, errSQLOpen error) {

	dbType := os.Getenv("DB_TYPE")
	if dbType == "postgres" {
		glog.Info("DB_TYPE: ", dbType)
		user := os.Getenv("POSTGRES_ENV_POSTGRES_USER")
		password := os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_ENV_DB_NAME")
		tcpAddr := os.Getenv("POSTGRES_ENV_PORT_5432_TCP_ADDR")
		sslMode := os.Getenv("POSTGRES_ENV_SSL_MODE")
		connString := fmt.Sprintf("postgres://" + user + ":" + password + "@" + tcpAddr + "/" + dbName + "?sslmode=" + sslMode)
		database, errSQLOpen = sql.Open("postgres", connString)
		if errSQLOpen != nil {
			return nil, errSQLOpen
		}
		statementPrepareCreateTable, errPrepareCreateTable := database.Prepare(postgresCreateTable)
		if errPrepareCreateTable != nil {
			return nil, errPrepareCreateTable
		}

		_, errExecCreateTable := statementPrepareCreateTable.Exec()
		if errExecCreateTable != nil {
			return nil, errExecCreateTable
		}
	} else if dbType == "mysql" {
		// user := os.Getenv("MYSQL_ENV_MYSQL_USER")
		// password := os.Getenv("MYSQL_ENV_MYSQL_PASSWORD")
		// dbName := os.Getenv("MYSQL_ENV_DB_NAME")
		// tcpAddr := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
		return nil, errors.New("unsupported yet")
	} else {
		database, errSQLOpen = sql.Open("sqlite3", "./db/users.db")
		if errSQLOpen != nil {
			return nil, errSQLOpen
		}
		statementPrepareCreateTable, errPrepareCreateTable := database.Prepare(sqlite3CreateTable)
		if errPrepareCreateTable != nil {
			return nil, errPrepareCreateTable
		}

		_, errExecCreateTable := statementPrepareCreateTable.Exec()
		if errExecCreateTable != nil {
			return nil, errExecCreateTable
		}
	}

	glog.Info("DB Ensured")
	return database, nil
}
