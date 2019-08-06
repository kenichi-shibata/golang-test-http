package utils

import (
	"database/sql"

	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
)

// ENV (POSTGRES OR MYSQL)
// defaults to sqlite for local dev running as volume mounted

func SQLOpen() (database *sql.DB, errSQLOpen error) {
	database, errSQLOpen = sql.Open("sqlite3", "./db/users.db")
	if errSQLOpen != nil {
		return nil, errSQLOpen
	}

	statementPrepareCreateTable, errPrepareCreateTable := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate TEXT)")
	if errPrepareCreateTable != nil {
		return nil, errPrepareCreateTable
	}

	_, errExecCreateTable := statementPrepareCreateTable.Exec()
	if errExecCreateTable != nil {
		return nil, errExecCreateTable
	}

	glog.Info("DB Ensured")
	return database, nil
}
