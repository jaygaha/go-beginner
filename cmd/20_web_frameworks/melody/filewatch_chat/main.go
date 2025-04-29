package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/olahol/melody"
)

/*
Cross-platform filesystem notifications for Go.
*/

func main() {
	file := "file.txt"

	// Create a new watcher.
	m := melody.New()
	w, _ := fsnotify.NewWatcher()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.tmpl")
	})

	// ws handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	// read file
	m.HandleConnect(func(s *melody.Session) {
		content, _ := os.ReadFile(file)

		s.Write(content)
	})

	// watch file
	go func() {
		for {
			select {
			case event := <-w.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					content, _ := os.ReadFile(file)
					m.Broadcast(content)
				}
			}
		}
	}()
	// add file to watcher
	w.Add(file)

	// start server
	fmt.Println("Server started at :8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		panic(err)
	}
}
