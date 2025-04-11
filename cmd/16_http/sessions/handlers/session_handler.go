package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStore is a global store for session management
// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
var SessionStore = sessions.NewCookieStore([]byte("secret-key"))

// init configures the session store
func init() {
	// set the session store options
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15, // 15 minutes
		HttpOnly: true,
		Secure:   false, // set to true if using HTTPS
	}
}

// Login handles the login request and init a new session
func Login(w http.ResponseWriter, r *http.Request) {
	// get the session
	session, _ := SessionStore.Get(r, "user-session")

	// set the session values
	session.Values["username"] = "hoge"
	session.Values["authenticated"] = true

	// save the session
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to html
	w.Header().Set("Content-Type", "text/html")
	// print message
	fmt.Fprintf(w, "Login successful!")
	fmt.Fprintf(w, "\n<a href=\"/user/profile\">My Profile</a>")
	fmt.Fprintf(w, "\n<a href=\"/auth/logout\">Logout</a>")

	return
}

// Profile handles the profile request
func Profile(w http.ResponseWriter, r *http.Request) {
	// get the session
	session, _ := SessionStore.Get(r, "user-session")

	// check if the user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	username, _ := session.Values["username"].(string)

	// print message
	fmt.Fprintf(w, "Profile page!")
	fmt.Fprintf(w, "Welcome, %s!", username)

	fmt.Fprintf(w, "\n<a href=\"/auth/logout\">Logout</a>")
	return
}

// Logout handles the logout request
func Logout(w http.ResponseWriter, r *http.Request) {
	// get the session
	session, _ := SessionStore.Get(r, "user-session")

	// revoke user authentication
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1 // delete session

	// save the session
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set content type to html
	w.Header().Set("Content-Type", "text/html")
	// print message
	fmt.Fprintf(w, "Logout successful!")
	fmt.Fprintf(w, "\n<a href=\"/auth/login\">Login</a>")
	fmt.Fprintf(w, "\t<a href=\"/\">Home</a>")
	return
}
