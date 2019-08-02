package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/golang/glog"
	"github.com/kenichi-shibata/golang-test-http/data"
	"github.com/kenichi-shibata/golang-test-http/utils"
)

const JsonTemplate = `{"message": "Hello {{.Username}}! Your birthday is in {{.DaysBeforeBirthday}} day(s)"}`
const JsonTemplate2 = `{"message": "Hello {{.Username}}! Happy Birthday!"}`

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/hello") {
		http.Error(w, "404 Not Found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	usernameInPath := strings.Replace(r.URL.Path, "/hello/", "", -1)

	if usernameInPath == "" {
		glog.Warning("Please input username")
		fmt.Fprintf(w, "{\"message\": \"Please input username\"}")
	} else {
		u := utils.User{Username: usernameInPath, DaysBeforeBirthday: 1, Birthdate: "2000-08-02"}

		switch r.Method {
		case "GET":
			glog.Info("GET")
		case "PUT":
			glog.Info("PUT")
		default:
			glog.Warning("Sorry, only GET and PUT methods are supported.")
		}

		errInsertDB := data.InsertDB(&u)
		if errInsertDB != nil {
			glog.Fatal(errInsertDB)
		}

		uCalc := data.SelectDB(&u)

		tmpl := template.New("User Template")
		var errTmplParse error

		if uCalc.DaysBeforeBirthday == 0 {
			tmpl, errTmplParse = tmpl.Parse(JsonTemplate2)
			if errTmplParse != nil {
				glog.Fatal("Parse: ", errTmplParse)
				return
			}
		} else {
			tmpl, errTmplParse = tmpl.Parse(JsonTemplate)
			if errTmplParse != nil {
				glog.Fatal("Parse: ", errTmplParse)
				return
			}
		}

		errTmplExecute := tmpl.Execute(w, uCalc)
		if errTmplExecute != nil {
			glog.Fatal("Execute: ", errTmplExecute)
			return
		}
	}
}

func main() {
	flag.Set("logtostderr", "false")
	flag.Set("stderrthreshold", "INFO")
	flag.Set("v", "2")
	flag.Parse()

	errSetupDB := data.SetupDB()
	if errSetupDB != nil {
		glog.Fatal(errSetupDB)
	}

	http.HandleFunc("/hello/", MainHandler)
	glog.Fatal(http.ListenAndServe(":8080", nil))
}
