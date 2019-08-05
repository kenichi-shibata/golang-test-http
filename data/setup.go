package data

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

const layoutISO = "2006-01-02" // https://golang.org/pkg/time/#Parse

func SetupDB() error {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		return errSQLOpen
	}

	statementPrepareCreateTable, errPrepareCreateTable := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate TEXT)")
	if errPrepareCreateTable != nil {
		return errPrepareCreateTable
	}

	_, errExecCreateTable := statementPrepareCreateTable.Exec()
	if errExecCreateTable != nil {
		return errExecCreateTable
	}

	glog.Info("DB Ensured")
	return nil
}

func InsertDB(user *utils.User) error {
	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		return errSQLOpen
	}

	statementPrepareInsertData, errPrepareInsertData := database.Prepare("INSERT INTO users (name, birthdate) VALUES (?, ?)")
	if errPrepareInsertData != nil {
		return errPrepareInsertData
	}

	_, errExecInsertData := statementPrepareInsertData.Exec(user.Username, user.Birthdate)
	if errExecInsertData != nil {
		return errExecInsertData
	}

	glog.Info("Insert Ensured")
	return nil
}

func SelectDB(user *utils.User) (userCalc utils.User, errSelectDB error) {
	var id int
	var name string
	var birthdate string

	database, errSQLOpen := sql.Open("sqlite3", "./users.db")
	if errSQLOpen != nil {
		glog.Fatal(errSQLOpen)
		panic(errSQLOpen)
	}
	defer database.Close()
	// this query needs to be changed to only return 1 or 0 entries not more than 1
	selectQuery := "SELECT id, name, birthdate FROM users WHERE name=?"
	rows := database.QueryRow(selectQuery, user.Username)
	errQuery := rows.Scan(&id, &name, &birthdate)
	glog.Info(errQuery)
	glog.Info("selectquery: ", selectQuery)

	switch {
	case errQuery == sql.ErrNoRows:
		glog.Warning("no user with username: ", user.Username)
		return utils.User{Username: user.Username, Birthdate: user.Birthdate, DaysBeforeBirthday: -365}, &utils.SQLRecordNotFoundError{Record: user.Username}
	case errQuery != nil:
		glog.Error("query error: ", errQuery)
		return utils.User{Username: user.Username, Birthdate: user.Birthdate, DaysBeforeBirthday: -365}, errQuery
	default:
		glog.Info("username is: ", user.Username)
		glog.Info("Found in DB: " + strconv.Itoa(id) + ": " + name + " " + birthdate)
		uCalc := calcDayDiff(birthdate, user)
		return uCalc, nil
	}
	// By default int falsy is 0 which makes it seem like its birthday day, if we leave it utils.User{}
	// return utils.User{Username: user.Username, Birthdate: user.Birthdate, DaysBeforeBirthday: -365}, errors.New("Unknown switch case")
}

func calcDayDiff(birthdate string, user *utils.User) (uCalc utils.User) {
	monthDayArray := strings.Split(birthdate, "-")[1:]

	birthdateWithYearSetToCurrent := strings.Join(append([]string{}, strconv.Itoa(time.Now().Year()), monthDayArray[0], monthDayArray[1]), "-")
	birthdateWithYearSetToCurrentParse, errBirthdateWithYearSetToCurrentParse := time.Parse(layoutISO, birthdateWithYearSetToCurrent)
	if errBirthdateWithYearSetToCurrentParse != nil {
		panic(errBirthdateWithYearSetToCurrentParse)
	}

	datetimeNow := time.Now()
	hourDiff := birthdateWithYearSetToCurrentParse.Sub(datetimeNow).Hours()
	dayDiff := int(math.Ceil(hourDiff / 24))
	// add another year if birthday already passed this year then add one year to birthdateWithYearSetToCurrent then call it birthdateWithYearSetToNext
	if dayDiff < 0 {
		birthdateWithYearSetToNext := birthdateWithYearSetToCurrentParse.AddDate(1, 0, 0) // add one year
		hourDiff = birthdateWithYearSetToNext.Sub(datetimeNow).Hours()
		dayDiff = int(math.Floor(hourDiff / 24))
	}

	glog.Info("birthdateWithYearSetToCurrent \t" + birthdateWithYearSetToCurrent)
	glog.Info("dateTimeNow \t\t\t\t" + datetimeNow.Format(layoutISO))
	glog.Info("dayDiff " + strconv.Itoa(dayDiff))
	glog.Info("Your birthday is " + strconv.Itoa(dayDiff) + " days from today!")

	uCalc = utils.User{Username: user.Username, Birthdate: user.Birthdate, DaysBeforeBirthday: dayDiff}
	return uCalc
}
