package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/olahol/melody"
)

var idCounter atomic.Int64 // atomic int64 type:

func main() {
	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.tmpl")
	})

	// serve ws
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// handle messages
	m.HandleConnect(func(s *melody.Session) {
		id := idCounter.Add(1)

		s.Set("id", id)
		s.Write(fmt.Appendf(nil, "iam %d", id))
	})

	// handle disconnect
	m.HandleDisconnect(func(s *melody.Session) {
		if id, ok := s.Get("id"); ok {
			m.BroadcastOthers(fmt.Appendf(nil, "user %d disconnected", id), s)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if id, ok := s.Get("id"); ok {
			m.BroadcastOthers(fmt.Appendf(nil, "user %d: %s", id, msg), s)
		}
	})

	// start server
	fmt.Println("gophy started on :8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
