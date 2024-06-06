package httptools

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// Server interface with a Start method
type Server interface {
    Start()
}

type server struct {
    httpServer *http.Server
}

func (s server) Start() {
    go func() {
        log.Println("Starting the HTTP server...") // Print a message that the server is starting
        err := s.httpServer.ListenAndServe() // Start the server
        log.Fatalf("HTTP server finished: %s. Finishing the process.", err) // Print an error message if the server stops
    }()
}

// CreateServer function makes a new Server
func CreateServer(port int, handler http.Handler) Server {
    return server{
        httpServer: &http.Server{
            Addr:           fmt.Sprintf(":%d", port), // Set the server address
            Handler:        handler, // Set the request handler
            ReadTimeout:    10 * time.Second, // Set the read timeout to 10 seconds
            WriteTimeout:   10 * time.Second, // Set the write timeout to 10 seconds
            MaxHeaderBytes: 1 << 20, // Set the max header size
        },
    }
}
