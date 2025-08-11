package auth

import (
	"fmt"
	"net/http"
)

func IsAuthenticated() bool {
	return false
}

func Login() {

	fmt.Println("User log in")

}

func Register() {
	fmt.Println("User Register")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User logout")
}
