package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Bag represents a simple bag structure.
type Bag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	SKU  int    `json:"sku"`
}

// in-memory storage
var (
	bags   = make(map[int]Bag) // map[ID]Bag
	nextId = 1
	mu     sync.Mutex // Mutex for thread safety
)

func CreateBag(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex before accessing bags
	// it helps to prevent concurrent access to bags
	mu.Lock()
	// Unlock the mutex after accessing bags
	defer mu.Unlock()

	var bag Bag
	if err := json.NewDecoder(r.Body).Decode(&bag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bag.ID = nextId
	bags[nextId] = bag
	nextId++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bag)
}

func GetBags(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	list := make([]Bag, 0, len(bags))
	for _, bag := range bags {
		list = append(list, bag)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(list) == 0 {
		json.NewEncoder(w).Encode(make([]Bag, 0))
		return
	}
	json.NewEncoder(w).Encode(list)
}

// get bag by id passed as a path parameter
func GetBag(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// path parameter
	id, err := extractID(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	bag, err := getBagByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bag)
}

// update bag by id passed as a path parameter
func UpdateBag(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := extractID(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	bag, err := getBagByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&bag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bags[id] = bag
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bag)
}

// delete bag by id passed as a path parameter
func DeleteBag(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := extractID(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	_, err = getBagByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	delete(bags, id)

	w.WriteHeader(http.StatusNoContent)
}

func extractID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getBagByID(id int) (Bag, error) {
	// directly access bag by ID from the map
	bag, exists := bags[id]
	if !exists {
		return Bag{}, fmt.Errorf("Bag not found")
	}
	return bag, nil
}
