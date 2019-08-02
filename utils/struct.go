package utils

type User struct {
	Username           string `json:"-"`           // This is coming from path not body
	DaysBeforeBirthday int    `json:"-"`           // This is not stored on DB or added via body but calculated on the fly
	Birthdate          string `json:"dateOfBirth"` // using layoutISO "2006-01-02" ref: https://stackoverflow.com/questions/20530327/origin-of-mon-jan-2-150405-mst-2006-in-golang
}

// Users DB Schema Mapping
// struct Username: DB name (path)
// struct Birthdate: DB birthdate (body)
// struct DaysBeforeBirthday: DB (nil)
