package auth

import (
	"fmt"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User logout")
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)

	Mutex.Lock()
	delete(Sessions, cookie.Value)
	Mutex.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
