package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Printf("Logging in the user with username %s and password %s\n", username, password)
}

