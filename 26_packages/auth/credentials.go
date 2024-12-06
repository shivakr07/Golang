package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}

//if you use small letter
// [for var, func, struct, ..., package] then it get scoped [kind of private] to that package only
//ex :- loginWithCredentials

// while if you make it capital then it visible [expose] out of this package
//gets exported automatically
