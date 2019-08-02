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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usernameInPath := strings.Replace(r.URL.Path, "/username/", "", -1)

	if usernameInPath == "" {
		glog.Warning("Please input username")
		fmt.Fprintf(w, "{\"message\": \"Please input username\"}")
	} else {
		u := utils.User{Username: usernameInPath, DaysBeforeBirthday: 1}

		tmpl := template.New("User Template")
		var errTmplParse error

		if u.DaysBeforeBirthday == 0 {
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

		errTmplExecute := tmpl.Execute(w, u)
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
	data.SetupDB()
	http.HandleFunc("/username/", handler)
	glog.Fatal(http.ListenAndServe(":8080", nil))
}
