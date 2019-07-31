// package data
package main

import (
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const layoutISO = "2006-01-02"

// func SetupDB() {
func main() {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		panic(errSQLOpen)
	}

	statementPrepareCreateTable, errPrepareCreateTable := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate DATE)")
	if errPrepareCreateTable != nil {
		panic(errPrepareCreateTable)
	}

	_, errExecCreateTable := statementPrepareCreateTable.Exec()
	if errExecCreateTable != nil {
		panic(errExecCreateTable)
	}

	statementPrepareInsertData, errPrepareInsertData := database.Prepare("INSERT INTO users (name, birthdate) VALUES (?, ?)")
	if errPrepareInsertData != nil {
		panic(errPrepareInsertData)
	}

	_, errExecInsertData := statementPrepareInsertData.Exec("Nic", "2019-08-31")
	if errExecInsertData != nil {
		panic(errExecInsertData)
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
		fmt.Println(strconv.Itoa(id) + ": " + name + " " + birthdate)
		monthDayArray := strings.Split(birthdate, "-")[1:]
		birthdateWithYearSetToCurrent := strings.Join(append([]string{}, strconv.Itoa(time.Now().Year()), monthDayArray[0], monthDayArray[1]), "-")
		birthdateWithYearSetToCurrentParse, errBirthdateWithYearSetToCurrentParse := time.Parse(layoutISO, birthdateWithYearSetToCurrent)
		if errBirthdateWithYearSetToCurrentParse != nil {
			panic(errBirthdateWithYearSetToCurrentParse)
		}
		datetimeNow := time.Now()
		hourDiff := birthdateWithYearSetToCurrentParse.Sub(datetimeNow).Hours()
		dayDiff := int(math.Round(hourDiff / 24))
		// fmt.Println("Your birthday is " + days + " from today!")
		fmt.Println("birthdateWithYearSetToCurrent " + birthdateWithYearSetToCurrent)
		fmt.Println("dayDiff " + strconv.Itoa(dayDiff))
	}
}
