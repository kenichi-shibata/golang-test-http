package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type User struct {
	Username           string
	DaysBeforeBirthday int
}

const JsonTemplate = `{"message": "Hello {{.Username}}! Your birthday is in {{.DaysBeforeBirthday}} day(s)"}`
const JsonTemplate2 = `{"message": "Hello {{.Username}}! Happy Birthday!"}`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usernameInPath := strings.Replace(r.URL.Path, "/username/", "", -1)

	if usernameInPath == "" {
		fmt.Fprintf(w, "{\"message\": \"Please input username\"}")
	} else {
		u := User{Username: usernameInPath, DaysBeforeBirthday: 1}

		tmpl := template.New("User Template")
		var errTmplParse error

		if u.DaysBeforeBirthday == 0 {
			tmpl, errTmplParse = tmpl.Parse(JsonTemplate2)
			if errTmplParse != nil {
				log.Fatal("Parse: ", errTmplParse)
				return
			}
		} else {
			tmpl, errTmplParse = tmpl.Parse(JsonTemplate)
			if errTmplParse != nil {
				log.Fatal("Parse: ", errTmplParse)
				return
			}
		}

		errTmplExectute := tmpl.Execute(w, u)
		if errTmplExectute != nil {
			log.Fatal("Execute: ", errTmplExectute)
			return
		}
	}
}

func main() {
	http.HandleFunc("/username/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
