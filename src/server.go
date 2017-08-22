package main

import (
	"log"
	"net/http"
	"sonMode"
)

//Open a https server on port 8080, which enable users/managers/super-managers do anything authorized.
// Welcome to NeverLand-Bank!
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()

	http.HandleFunc("/", sonMode.Welcome)
	http.HandleFunc("/signIn", sonMode.SignInHandle)
	http.HandleFunc("/managerSignIn", sonMode.ManagerSignInHandle)
	http.HandleFunc("/superManagerSignIn", sonMode.SuperManagerSignInHandle)
	http.HandleFunc("/signOut", sonMode.SignOutHandle)
	http.HandleFunc("/signUp", sonMode.SignUpHandle)
	http.HandleFunc("/homepage", sonMode.HomepageHandle)
	http.HandleFunc("/profile", sonMode.ProfileHandle)
	http.HandleFunc("/transfer", sonMode.TransferHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
