package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var target = flag.String("target", "http://localhost:8090", "request target")

func main() {
	flag.Parse()
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout
	}
	// Loop every second
	for range time.Tick(1 * time.Second) {
		// Send a GET request to the target URL
		resp, err := client.Get(fmt.Sprintf("%s/api/v1/some-data", *target))
		if err == nil {
			// If no error, print the response status code
			log.Printf("response %d", resp.StatusCode)
		} else {
			// If there's an error, print the error message
			log.Printf("error %s", err)
		}
	}
}
