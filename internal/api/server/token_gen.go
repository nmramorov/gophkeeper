package server

func GenerateToken(login, password string) string {
	return login + password + "salt"
}
