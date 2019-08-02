package data

// package main

import (
	"database/sql"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/kenichi-shibata/golang-test-http/utils"
	_ "github.com/mattn/go-sqlite3"
)

const layoutISO = "2006-01-02"

func SetupDB() {
	// func main() {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		panic(errSQLOpen)
	}

	statementPrepareCreateTable, errPrepareCreateTable := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate TEXT)")
	if errPrepareCreateTable != nil {
		panic(errPrepareCreateTable)
	}

	_, errExecCreateTable := statementPrepareCreateTable.Exec()
	if errExecCreateTable != nil {
		panic(errExecCreateTable)
	}

	glog.Info("DB Ensured")
}

func InsertDB(user *utils.User) {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		glog.Fatal(errSQLOpen)
		panic(errSQLOpen)
	}

	statementPrepareInsertData, errPrepareInsertData := database.Prepare("INSERT INTO users (name, birthdate) VALUES (?, ?)")
	if errPrepareInsertData != nil {
		glog.Fatal(errPrepareInsertData)
		panic(errPrepareInsertData)
	}

	execInsertData, errExecInsertData := statementPrepareInsertData.Exec(user.Username, user.Birthdate)
	if errExecInsertData != nil {
		glog.Fatal(errExecInsertData)
		panic(errExecInsertData)
	}

	glog.Info("Insert Ensured")
	glog.Info(execInsertData)
}

func SelectDB(user *utils.User) {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		glog.Fatal(errSQLOpen)
		panic(errSQLOpen)
	}

	rows, _ := database.Query("SELECT id, name, birthdate FROM users")

	var id int
	var name string
	var birthdate string
	for rows.Next() {
		err := rows.Scan(&id, &name, &birthdate)
		if err != nil {
			panic(err)
		}
		glog.Info("Found in DB: " + strconv.Itoa(id) + ": " + name + " " + birthdate)
		monthDayArray := strings.Split(birthdate, "-")[1:]
		birthdateWithYearSetToCurrent := strings.Join(append([]string{}, strconv.Itoa(time.Now().Year()), monthDayArray[0], monthDayArray[1]), "-")
		birthdateWithYearSetToCurrentParse, errBirthdateWithYearSetToCurrentParse := time.Parse(layoutISO, birthdateWithYearSetToCurrent)
		if errBirthdateWithYearSetToCurrentParse != nil {
			panic(errBirthdateWithYearSetToCurrentParse)
		}
		datetimeNow := time.Now()
		hourDiff := birthdateWithYearSetToCurrentParse.Sub(datetimeNow).Hours()
		dayDiff := int(math.Round(hourDiff / 24))
		// add another year if birthday already passed this year then add one year to birthdateWithYearSetToCurrent then call it birthdateWithYearSetToNext
		if dayDiff < 0 {
			birthdateWithYearSetToNext := birthdateWithYearSetToCurrentParse.AddDate(1, 0, 0) // add one year
			hourDiff = birthdateWithYearSetToNext.Sub(datetimeNow).Hours()
			dayDiff = int(math.Round(hourDiff / 24))
		}
		glog.Info("birthdateWithYearSetToCurrent \t" + birthdateWithYearSetToCurrent)
		glog.Info("dateTimeNow \t\t\t\t" + datetimeNow.Format(layoutISO))
		glog.Info("dayDiff " + strconv.Itoa(dayDiff))
		glog.Info("Your birthday is " + strconv.Itoa(dayDiff) + " days from today!")
	}
}
