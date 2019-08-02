package utils

type User struct {
	Username           string
	DaysBeforeBirthday int // this is not stored on DB but calculated on the fly
	Birthdate          string
}
