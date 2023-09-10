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

type BinaryData struct {
	UUID   string
	UserID string
	Data   []byte
	Meta   string
}

type BankCardData struct {
	UUID       string
	UserID     string
	Number     string
	Owner      string
	ExpiresAt  string
	SecretCode string
	PinCode    string
	Meta       string
}
