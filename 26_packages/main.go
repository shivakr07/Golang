package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shivakr07/authbuddy/auth"
	"github.com/shivakr07/authbuddy/user"
)

func main() {

	//here we are using the auth package
	auth.LoginWithCredentials("shiva", "secret")

	session := auth.GetSession()
	fmt.Println(session)

	user1 := user.User{
		Email: "user@email.com",
		Name:  "Alan Walker",
	}

	fmt.Println(user1.Email, " ", user1.Name)
	color.Red(user1.Email)
	color.Green(user1.Name)

}

//to create a packagewe generally use a very nice convention
// first we initiate a module [our project]
// go mod init github.com/shivakr07/authbuddy
//go init modulename

//-> it gives you a go.mod file
// and then you can use that package

// package means a folder have file [files]

//--------------THIRD PARTY PACKAGE---------------------------
// to install third party packages we use
// go get github.com/fatih/color

//go sum is like package.json which maintains versioning of the package
