package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/centrifugal/centrifuge"
)

// setupNode initializes the Centrifuge node
func setupNode() *centrifuge.Node {
	// Configure node with logging
	cfg := centrifuge.Config{
		LogLevel: centrifuge.LogLevelDebug,
		LogHandler: func(entry centrifuge.LogEntry) {
			log.Printf("%s: %v", entry.Message, entry.Fields)
		},
	}

	node, err := centrifuge.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	// Handle connected clients
	node.OnConnect(func(client *centrifuge.Client) {

		// Handle subscriptions
		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			if e.Channel != "chat" {
				log.Printf("Subscription error: invalid channel %s", e.Channel)
				cb(centrifuge.SubscribeReply{}, &centrifuge.Error{Code: 100, Message: "Only 'chat' channel allowed"})
				return
			}
			log.Printf("Client %s subscribed to %s", client.ID(), e.Channel)
			cb(centrifuge.SubscribeReply{}, nil)
		})

		// Handle published messages
		client.OnPublish(func(e centrifuge.PublishEvent, cb centrifuge.PublishCallback) {
			// Validate JSON payload
			var payload map[string]interface{}
			if err := json.Unmarshal(e.Data, &payload); err != nil {
				log.Printf("Publish error: invalid JSON: %v", err)
				cb(centrifuge.PublishReply{}, &centrifuge.Error{Code: 3501, Message: "bad request: invalid JSON"})
				return
			}
			if _, exists := payload["input"]; !exists {
				log.Printf("Publish error: missing 'input' field")
				cb(centrifuge.PublishReply{}, &centrifuge.Error{Code: 3501, Message: "bad request: missing 'input' field"})
				return
			}
			log.Printf("Message published to %s: %s", e.Channel, string(e.Data))
			cb(centrifuge.PublishReply{}, nil)
		})

		// Handle disconnections
		client.OnDisconnect(func(e centrifuge.DisconnectEvent) {
			log.Printf("Client %s disconnected: code=%d, reason=%s", client.ID(), e.Code, e.Reason)
		})
	})

	return node
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		newCtx := centrifuge.SetCredentials(ctx, &centrifuge.Credentials{
			UserID:   "42",
			ExpireAt: time.Now().Unix() + 60,
			Info:     []byte(`{"name": "Alexander"}`),
		})
		r = r.WithContext(newCtx)
		h.ServeHTTP(w, r)
	})
}

func main() {
	node := setupNode()

	// Start the node
	if err := node.Run(); err != nil {
		log.Fatalf("Failed to run node: %v", err)
	}

	mux := http.DefaultServeMux
	// Set up WebSocket handler
	wsHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{
		ReadBufferSize:     1024,
		UseWriteBufferPool: true,
	})

	mux.Handle("/connection/websocket", authMiddleware(wsHandler))

	// Serve static files
	mux.Handle("/", http.FileServer(http.Dir("client")))

	log.Println("Server running on :8800")
	if err := http.ListenAndServe(":8800", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
