package auth

import (
	"fmt"
	"forum/database"
	"forum/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type session struct {
	Username string
	Expiry   time.Time
}

var Mutex sync.RWMutex
var Sessions = map[string]session{}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	Exist, err := database.Login(username, password)
	if err != nil {
		panic(err)
	}

	if !Exist {
		utils.ToastError(w, "User not found", "#login")
		fmt.Println("User not found:", username)
		return
	}

	uuidV4, err := uuid.NewV4()
	if err != nil {
		utils.ErrorHandler(w, "Internal server error", 500)
		return
	}
	uuid := uuidV4.String()
	expiresAt := time.Now().Add(24 * time.Hour)

	Mutex.Lock()
	Sessions[uuid] = session{
		Username: username,
		Expiry:   expiresAt,
	}
	Mutex.Unlock()

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    uuid,
		HttpOnly: true,  // Non lisible par JavaScript (sécurité)
		Secure:   false, // Mettre true si HTTPS (true on production)
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	fmt.Println("User logged in", username)
}

func GetSession(r *http.Request) (session, bool) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return session{}, false
	}

	Mutex.RLock()
	sess, ok := Sessions[cookie.Value]
	Mutex.RUnlock()

	if !ok || sess.Expiry.Before(time.Now()) {
		return session{}, false
	}

	return sess, true
}
