package main

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReport_Process(t *testing.T) {
	// Creating a test request
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("lb-author", "test-author") // Setting the lb-author header
	req.Header.Set("lb-req-cnt", "1")          // Setting the lb-req-cnt header

	// Creating an instance of the Report structure
	r := make(Report)

	// Testing the Process method for adding data to the report
	r.Process(req)
	if !reflect.DeepEqual(r["test-author"], []string{"1"}) {
		t.Errorf("Unexpected report state %s", r)
	}

	// Changing the value of the lb-req-cnt header
	req.Header.Set("lb-req-cnt", "2")
	// Retesting the Process method for adding data to the report
	r.Process(req)
	if !reflect.DeepEqual(r["test-author"], []string{"1", "2"}) {
		t.Errorf("Unexpected report state %s", r)
	}

	// Changing the value of the lb-author header
	req.Header.Set("lb-author", "test-len")
	// Retesting the Process method with a large number of requests
	for i := 0; i < 103; i++ {
		req.Header.Set("lb-req-cnt", "test-len")
		r.Process(req)
	}
	// Checking that the length of the array is correct
	if len(r["test-len"]) != reportMaxLen {
		t.Errorf("Unexpected error length: %d", len(r["test-len"]))
	}
}
