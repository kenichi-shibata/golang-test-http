package utils

type User struct {
	Username           string
	DaysBeforeBirthday int    // this is not stored on DB but calculated on the fly
	Birthdate          string // using layoutISO "2006-01-02" ref: https://stackoverflow.com/questions/20530327/origin-of-mon-jan-2-150405-mst-2006-in-golang
}

// Users DB Schema Mapping
// struct Username: DB name
// struct Birthdate: DB birthdate
// struct DaysBeforeBirthday: DB (nil)
