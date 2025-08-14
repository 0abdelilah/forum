package auth

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type session struct {
	Id       int
	Username string
	Expiry   time.Time
}

var Mutex sync.RWMutex
var Sessions = map[string]session{}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		User string `json:"username"`
		Pass string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid JSON",
		})
		return
	}

	username := creds.User
	password := creds.Pass

	exist, err := database.Login(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Internal server error",
		})
		fmt.Println(err)
		return
	}

	if !exist {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Invalid username or password",
		})
		return
	}

	uuidV4, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"error":   "Internal server error",
		})
		return
	}
	SessUUID := "sess_" + uuidV4.String()
	expiresAt := time.Now().Add(24 * time.Hour)

	Mutex.Lock()
	Sessions[SessUUID] = session{
		Username: username,
		Expiry:   expiresAt,
	}
	Mutex.Unlock()

	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   SessUUID,
		Expires: expiresAt,
		Path:    "/",
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(map[string]string{
		"success": "true",
	})
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
