package handlers

import (
	"fmt"
	"net/http"
	"time"
)

/*
Setting a Cookie for the client:

Setting up using default cookie handler:
  - http.SetCookie(w, &http.Cookie{
    Name: "cookie-name",
    Value: "cookie-value",
    Path: "/",
    Domain: "localhost",
    Expires: time.Now().Add(365 * 24 * time.Hour),
    Secure: true,
    HttpOnly: true,
    SameSite: http.SameSiteStrictMode,
    })
*/
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "go-cookie",
		Value:    "go-cookie-value",
		Expires:  time.Now().Add(24 * time.Hour), // 1 day
		Path:     "/",                            // all paths
		HttpOnly: true,                           // only accessible via http,
		Secure:   false,                          // only accessible via https
	}

	http.SetCookie(w, &cookie) // set cookie

	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	// print message
	fmt.Fprintf(w, "Cookie set!")
	fmt.Fprintf(w, "\n<a href=\"/\">Back</a>")

	return
}

/*
Getting a Cookie from the client:

Getting using default cookie handler:
  - cookie, err := r.Cookie("cookie-name")
  - if err != nil {
    // handle error
    }
*/
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("go-cookie")

	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	if err != nil {
		switch err {
		case http.ErrNoCookie:
			fmt.Fprintf(w, "Cookie not found")
		default:
			fmt.Fprintf(w, "Error reading cookie: %v", err)
		}

		fmt.Fprintf(w, "\n<a href=\"/\">Back</a>")
		return
	}

	fmt.Fprintf(w, "Cookie value: %s", cookie.Value)
	fmt.Fprintf(w, "\n<a href=\"/\">Back</a>")
}

/*
Deleting a Cookie from the client:

Deleting using default cookie handler:
  - cookie := http.Cookie{
    Name: "cookie-name",
    Value: "",
    Expires: time.Now(),
    Path: "/",
    }
  - http.SetCookie(w, &cookie)
*/

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "go-cookie",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour), // negative time
		Path:    "/",
	}

	http.SetCookie(w, &cookie)

	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "Cookie deleted")

	fmt.Fprintf(w, "\n<a href=\"/\">Back</a>")
}
