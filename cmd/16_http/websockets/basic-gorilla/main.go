package main

import (
	"flag"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

/*
Websocket
- it provide full-duplex communication channel over a single TCP, long-lived connection allowing for multiple messages to be sent and received at the same time
- it is a persistent connection, which means that the connection is maintained until the client or server decides to close it

Websocket vs HTTP and other protocols
- HTTP is a stateless protocol, which means that each request is independent of the previous one.
- This model is inconvenient for real-time applications, which require a persistent connection.

- Websocket is a persistent connection, which means that the connection is maintained until the client or server decides to close it.
- This reduces latency and increases performance.

- Server SentEvents (SSE) is another protocol for real-time updates, where the server pushes updates to the client over an HTTP connection; however SSE is
- is one way (sever to client) and not as flexible as Websocket.

Use cases:
- chat applications
- multiplayer games
- real time data feeds (stock prices, weather, etc.)
- live updates (e.g. news, updates, etc.)
- collaborative editing (e.g. Google Docs, etc.)
- etc.
*/

// template handler; it defines a handler that will be used to serve templates
type templateHandler struct {
	once     sync.Once // once is used to ensure that the template is parsed only once
	fileName string
	tmpl     *template.Template // tmpl is the template to be executed and *template.Template is a pointer to a template
}

// ServeHTTP serves HTTP requests
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// parse the template only once
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.fileName))) // filepath.Join concatenates the path segments with the OS specific separator
	})
	// execute the template
	t.tmpl.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8800", "The addr of the application.") // addr is the address of the application
	flag.Parse()                                                            // parse the flags

	r := newRoom() // create a new room

	// routes
	http.Handle("/", &templateHandler{fileName: "chat.tmpl"}) // handle the root path with the template handler
	http.Handle("/rooms", r)                                  // handle the rooms path with the room

	// get the room going
	go r.run()

	// start the web server
	err := http.ListenAndServe(*addr, nil) // listen and serve the application
	if err != nil {
		panic(err)
	}
}
