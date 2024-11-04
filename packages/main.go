package main

import (
	"auth/auth"
	"auth/user"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	auth.LoginWithCredentials("nitish", "nitish")
	session := auth.GetSession()

	fmt.Println(session)

	user := user.User{
		Email: "nitish", Username: "nanohard", Passwowrd: "ntitiit",
	}

	color.Green(user.Email)
}
