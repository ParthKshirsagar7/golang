package auth

func extractSession() string {
	return "Logged in"
}

func GetSession() string {
	return extractSession()
}