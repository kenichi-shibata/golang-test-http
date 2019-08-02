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

	// Test 4: Method POST not allowed
	req4, err4 := http.NewRequest("POST", "/hello/", nil)
	if err4 != nil {
		t.Fatal(err4)
	}
	rr4 := httptest.NewRecorder()
	handler4 := http.HandlerFunc(MainHandler)

	handler4.ServeHTTP(rr4, req4)

	if status := rr4.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
	// Test 5: Method PUT no user
	req5, err5 := http.NewRequest("PUT", "/hello/", nil)
	if err5 != nil {
		t.Fatal(err5)
	}
	rr5 := httptest.NewRecorder()
	handler5 := http.HandlerFunc(MainHandler)

	handler5.ServeHTTP(rr5, req5)

	if status := rr5.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected5 := `{"message": "Please input username"}`
	if rr5.Body.String() != expected5 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr5.Body.String(), expected5)
	}
	// Test 6: Method PUT happy path
	req6, err6 := http.NewRequest("PUT", "/hello/health-check", nil)
	if err6 != nil {
		t.Fatal(err6)
	}
	rr6 := httptest.NewRecorder()
	handler6 := http.HandlerFunc(MainHandler)

	handler6.ServeHTTP(rr6, req6)

	if status := rr6.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

}
