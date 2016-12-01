package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

/*
Expect Http Code: 200
Expect Content-Type: application/json
 */
func TestGetTasks(t *testing.T) {
	talkController := &TalkController{}

	req, err := http.NewRequest("GET", "/talks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(talkController.Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content type header does not match: got %v want %v",
		ctype, "application/json")
	}
}