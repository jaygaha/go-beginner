package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/olahol/melody"
)

// message struct; its a struct of user and msg
type Message struct {
	User      string    `json:"user"`
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// create a new melody instance
	m := melody.New()

	// ask user if initial session
	m.HandleConnect(func(s *melody.Session) {
		q := s.Request.URL.Query()
		user := q.Get("user")
		if user == "" {
			// append user timestamp
			time := strconv.FormatInt(time.Now().Unix(), 10)
			user = "user_" + time
		}

		// set user to session
		s.Set("user", user)
	})

	// handle message
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// get user from session
		user, exists := s.Get("user")

		if !exists {
			time := strconv.FormatInt(time.Now().Unix(), 10)
			user = "user_" + time
		}

		// create a new message
		message := Message{
			User:      user.(string),
			Msg:       string(msg),
			CreatedAt: time.Now(),
		}

		// JSON encode the message
		jsonMsg, err := json.Marshal(message)
		if err != nil {
			return
		}

		// broadcast the message
		err = m.Broadcast(jsonMsg)
		if err != nil {
			fmt.Println("Error broadcasting message: ", err)
		}
	})

	// handlue route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.tmpl")
	})

	// handle ws
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// start server
	fmt.Println("Websocket started on port 8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
