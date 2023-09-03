package models

type User struct {
	Login    string
	Password string
}

type CredentialsData struct {
	UUID     string
	Login    string
	Password string
	Meta     string
}
