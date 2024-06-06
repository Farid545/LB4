package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const reportMaxLen = 100 // Maximum length of report data for each author

// Report represents a map of string slices
type Report map[string][]string

// Process processes the incoming HTTP request and updates the report
func (r Report) Process(req *http.Request) {
	author := req.Header.Get("lb-author")    // Get the author from the request header
	counter := req.Header.Get("lb-req-cnt")  // Get the request counter from the request header
	log.Printf("GET some-data from [%s] request [%s]", author, counter) // Log the request details

	// Check if the author name is not empty
	if len(author) > 0 {
		list := r[author] // Get the list of requests for the author
		list = append(list, counter) // Append the new request counter to the list

		// Trim the list if its length exceeds the maximum allowed length
		if len(list) > reportMaxLen {
			list = list[len(list)-reportMaxLen:]
		}

		// Update the report with the trimmed list
		r[author] = list
	}
}

// ServeHTTP serves the report data as JSON over HTTP
func (r Report) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("content-type", "application/json") // Set the content type header to JSON
	rw.WriteHeader(http.StatusOK) // Set the HTTP status code to 200 (OK)

	// Encode the report data as JSON and write it to the response writer
	_ = json.NewEncoder(rw).Encode(r)
}
