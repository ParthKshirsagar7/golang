package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/parthkshirsagar7/golang/auth"
	"github.com/parthkshirsagar7/golang/user"
)

func main() {
	auth.LoginWithCredentials("admin", "admin")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "1@example.in",
		Name: "Test User",
	}

	color.Cyan(user.Email)
	color.Cyan(user.Name)
}
