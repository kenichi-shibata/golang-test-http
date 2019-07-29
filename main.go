package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type User struct {
	Username string
}

const JsonTemplate = `{"message": "Hello {{.Username}}! Your birthday is in N days"}`

func handler(w http.ResponseWriter, r *http.Request) {
	usernameInPath := strings.Replace(r.URL.Path, "/username/", "", -1)
	u := User{Username: usernameInPath}

	tmpl := template.New("User Template")

	tmpl, errTmplParse := tmpl.Parse(JsonTemplate)
	if errTmplParse != nil {
		log.Fatal("Parse: ", errTmplParse)
		return
	}

	errTmplExectute := tmpl.Execute(w, u)
	if errTmplExectute != nil {
		log.Fatal("Execute: ", errTmplExectute)
		return
	}
}

func main() {
	http.HandleFunc("/username/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
