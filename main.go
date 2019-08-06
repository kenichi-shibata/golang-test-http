package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

	placeholderBirthdate := "2000-08-02"
	usernameInPath := strings.Replace(r.URL.Path, "/hello/", "", -1)
	u := utils.User{Username: usernameInPath, DaysBeforeBirthday: 1, Birthdate: &placeholderBirthdate}

	switch r.Method {
	case "GET":
		glog.Info("GET")
		w.Header().Set("Content-Type", "application/json")

		if usernameInPath == "" {
			glog.Warning("Please input username")
			fmt.Fprintf(w, "{\"message\": \"Please input username\"}")
		} else {
			uCalc, errSelectDB := data.SelectDB(&u)
			if errSelectDB != nil {
				switch errSelectDB.(type) {
				case *utils.SQLRecordNotFoundError:
					glog.Warning("Select: ", errSelectDB)
					fmt.Fprintf(w, "{\"message\": \"Hello %v! We didn't find your birthday on our records! Please PUT your birthdate.\"}", usernameInPath)
					return
				default:
					glog.Error("Select: ", errSelectDB)
					return
				}
			}

			tmpl := template.New("User Template")
			var errTmplParse error

			if uCalc.DaysBeforeBirthday == 0 {
				tmpl, errTmplParse = tmpl.Parse(JsonTemplate2)
				if errTmplParse != nil {
					glog.Error("Parse: ", errTmplParse)
					return
				}
			} else {
				tmpl, errTmplParse = tmpl.Parse(JsonTemplate)
				if errTmplParse != nil {
					glog.Error("Parse: ", errTmplParse)
					return
				}
			}

			errTmplExecute := tmpl.Execute(w, uCalc)
			if errTmplExecute != nil {
				glog.Error("Execute: ", errTmplExecute)
				return
			}
		}
	case "PUT":
		glog.Info("PUT")
		w.Header().Set("Content-Type", "application/json")

		if usernameInPath == "" {
			glog.Warning("Please input username")
			fmt.Fprintf(w, "{\"message\": \"Please input username\"}")
			return
		} else {
			body, errReadBody := ioutil.ReadAll(r.Body)
			if errReadBody != nil {
				glog.Error("errReadBody:", errReadBody)
				http.Error(w, "Empty Body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			var userFromBody *utils.User
			errUnmarshalBody := json.Unmarshal(body, &userFromBody)
			if errUnmarshalBody != nil {
				glog.Warning("No JSON Found in Body:", errUnmarshalBody)
				http.Error(w, "No JSON Found in Body", 500)
				return
			}
			if userFromBody.Birthdate == nil {
				glog.Error("dateOfBirth field required")
				http.Error(w, "dateOfBirth Field required", 500)
				return
			}
			glog.Info("unmarshalled user: ", &userFromBody)

			insertUser := utils.User{Username: usernameInPath, Birthdate: userFromBody.Birthdate}
			errInsertDB := data.InsertDB(&insertUser)
			if errInsertDB != nil {
				glog.Error(errInsertDB)
				http.Error(w, errInsertDB.Error(), 500)
				return
			}
			w.WriteHeader(204)
		}
	default:
		glog.Warning("Only GET and PUT methods are supported Method Called::", r.Method)
		http.Error(w, "405 Method Not Allowed", 405)
		return
	}
}

func main() {
	flag.Set("logtostderr", "false")
	flag.Set("stderrthreshold", "INFO")
	flag.Set("v", "2")
	flag.Parse()

	http.HandleFunc("/hello/", MainHandler)
	glog.Fatal(http.ListenAndServe(":8080", nil))
}
