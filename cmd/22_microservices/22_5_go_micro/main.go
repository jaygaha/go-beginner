package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/web"
)

// Movie represents a movie in the rental service
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	PosterURL   string `json:"poster_url"`
	IsRented    bool   `json:"is_rented"`
}

// RentalService manages the in-memory movie database
type RentalService struct {
	movies map[string]*Movie
	mutex  sync.RWMutex
}

// NewRentalService initializes the service with sample movies
func NewRentalService() *RentalService {
	return &RentalService{
		movies: map[string]*Movie{
			"1":  {ID: "MS01", Title: "The Lord of the Rings", ReleaseYear: 1978, PosterURL: "https://image.tmdb.org/t/p/w200/liW0mjvTyLs7UCumaHhx3PpU4VT.jpg", IsRented: false},
			"2":  {ID: "MS02", Title: "The Lord of the Rings: The War of the Rohirrim", ReleaseYear: 2024, PosterURL: "https://image.tmdb.org/t/p/w200/cXzCOx1hUuU9CfmiEv6rXjb6EqU.jpg", IsRented: false},
			"3":  {ID: "MS03", Title: "The Lord of the Rings: The Fellowship of the Ring", ReleaseYear: 2001, PosterURL: "https://image.tmdb.org/t/p/w200/6oom5QYQ2yQTMJIbnvbkBL9cHo6.jpg", IsRented: false},
			"4":  {ID: "MS04", Title: "The Lord of the Rings: The Return of the King", ReleaseYear: 2003, PosterURL: "https://image.tmdb.org/t/p/w200/rCzpDGLbOoPwLjy3OAm5NUPOTrC.jpg", IsRented: false},
			"5":  {ID: "MS05", Title: "The Lord of the Rings: The Two Towers", ReleaseYear: 2002, PosterURL: "https://image.tmdb.org/t/p/w200/5VTN0pR8gcqV3EPUHHfMGnJYN9L.jpg", IsRented: false},
			"6":  {ID: "MS06", Title: "The Lord of the Rings: The Rings of Power Global Fan Screening", ReleaseYear: 2001, PosterURL: "https://image.tmdb.org/t/p/w200/yHA9Fc37VmpUA5UncTxxo3rTGVA.jpg", IsRented: false},
			"7":  {ID: "MS07", Title: "The Hobbit: An Unexpected Journey", ReleaseYear: 2012, PosterURL: "https://image.tmdb.org/t/p/w200/xT98tLqatZPQApyRmlPL12LtiWp.jpg", IsRented: false},
			"8":  {ID: "MS08", Title: "The Hobbit: The Battle of the Five Armies", ReleaseYear: 2014, PosterURL: "https://image.tmdb.org/t/p/w200/xT98tLqatZPQApyRmlPL12LtiWp.jpg", IsRented: false},
			"9":  {ID: "MS09", Title: "The Hobbit: The Desolation of Smaug", ReleaseYear: 2013, PosterURL: "https://image.tmdb.org/t/p/w200/xQYiXsheRCDBA39DOrmaw1aSpbk.jpg", IsRented: false},
			"10": {ID: "MS10", Title: "J. R. R. Tolkien's The Hobbit", ReleaseYear: 1967, PosterURL: "https://image.tmdb.org/t/p/w200/sm8h3bhCtVe7cGS7yzFuFoiO46G.jpg", IsRented: false},
		},
	}
}

// ListMovies handles GET /movies
func (rs *RentalService) ListMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rs.mutex.RLock()
	defer rs.mutex.RUnlock()

	var movies []Movie
	for _, movie := range rs.movies {
		movies = append(movies, *movie)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		logger.Errorf("Failed to encode movies: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("Listed movies")
}

// RentMovie handles POST /rent
func (rs *RentalService) RentMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		MovieID string `json:"movie_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("Invalid JSON: %v", err)
		http.Error(w, "Bad request: invalid JSON", http.StatusBadRequest)
		return
	}

	if req.MovieID == "" {
		logger.Error("Movie ID is required")
		http.Error(w, "Bad request: movie_id is required", http.StatusBadRequest)
		return
	}

	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	// find movie if exists
	logger.Infof("Renting movie: %s", rs.movies)
	movie, exists := findMovieByID(req.MovieID, rs.movies)
	// movie, exists := rs.movies[req.MovieID]
	if !exists {
		logger.Errorf("Movie not found: %s", req.MovieID)
		http.Error(w, "Not found: movie does not exist", http.StatusNotFound)
		return
	}

	if movie.IsRented {
		logger.Errorf("Movie already rented: %s", req.MovieID)
		http.Error(w, "Conflict: movie is already rented", http.StatusConflict)
		return
	}

	movie.IsRented = true
	resp := struct {
		Message string `json:"message"`
	}{Message: "Movie rented successfully"}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Infof("Rented movie: %s", movie.Title)
}

func withCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func findMovieByID(id string, movies map[string]*Movie) (*Movie, bool) {
	for _, movie := range movies {
		if movie.ID == id {
			return movie, true
		}
	}
	return nil, false
}

func main() {
	// Create a new Micro service
	// Create a web service for HTTP handling
	webService := web.NewService(
		web.Name("movie-rental.web"),
		web.Address(":8800"),
	)

	// Initialize the rental service
	rs := NewRentalService()

	// Register HTTP handlers
	webService.HandleFunc("/movies", withCORS(rs.ListMovies))
	webService.HandleFunc("/rent", withCORS(rs.RentMovie))

	// Initialize and run the web service
	if err := webService.Init(); err != nil {
		logger.Fatalf("Failed to initialize web service: %v", err)
	}

	// Run the service
	if err := webService.Run(); err != nil {
		logger.Fatalf("Failed to run web service: %v", err)
	}
}
