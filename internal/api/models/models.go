package models

type User struct {
	UUID     string
	Login    string
	Password string
	Token    string
}

type CredentialsData struct {
	UUID     string
	UserID   string
	Login    string
	Password string
	Meta     string
}

type TextData struct {
	UUID   string
	UserID string
	Data   string
	Meta   string
}
