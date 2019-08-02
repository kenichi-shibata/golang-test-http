package utils

type User struct {
	Username           string `json:"username"`
	DaysBeforeBirthday int    `json:"-"`         // this is not stored on DB or added via body but calculated on the fly
	Birthdate          string `json:"birthdate"` // using layoutISO "2006-01-02" ref: https://stackoverflow.com/questions/20530327/origin-of-mon-jan-2-150405-mst-2006-in-golang
}

// Users DB Schema Mapping
// struct Username: DB name
// struct Birthdate: DB birthdate
// struct DaysBeforeBirthday: DB (nil)
