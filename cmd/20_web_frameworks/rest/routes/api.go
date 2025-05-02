package routes

import (
	"net/http"

	"github.com/jaygaha/go_beginner/cmd/20_web_frameworks/rest/handlers"
)

func RegisterRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// if id is passed as query string, get bag by id, else get all bags
		if r.URL.Query().Get("id") != "" {
			handlers.GetBag(w, r)
		} else {
			handlers.GetBags(w, r)
		}
	case http.MethodPost:
		handlers.CreateBag(w, r)
	case http.MethodPut:
		handlers.UpdateBag(w, r)
	case http.MethodDelete:
		handlers.DeleteBag(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
