package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	// Test 1: wrong path /health-check
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	//  Test 2: Happy Path Happy Birthday/hello/health-check
	req2, err2 := http.NewRequest("GET", "/hello/health-check", nil)
	if err2 != nil {
		t.Fatal(err2)
	}
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(MainHandler)

	handler2.ServeHTTP(rr2, req2)

	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected2 := `{"message": "Hello health-check! Happy Birthday!"}`
	if rr2.Body.String() != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr2.Body.String(), expected2)
	}
	//  Test 3: No username /hello/
	req3, err3 := http.NewRequest("GET", "/hello/", nil)
	if err3 != nil {
		t.Fatal(err3)
	}
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(MainHandler)

	handler3.ServeHTTP(rr3, req3)

	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected3 := `{"message": "Please input username"}`
	if rr3.Body.String() != expected3 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr3.Body.String(), expected3)
	}
}
