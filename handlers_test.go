package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUsersIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	handler := ContentType(http.HandlerFunc(UsersIndex))

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got '%v' want '%v'",
			status, http.StatusOK)
	}

	if ctype := responseRecorder.Header().Get("Content-Type"); ctype != jsonHeader {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, jsonHeader)
	}

	expected := "[]"
	response := responseRecorder.Body.String()
	response = strings.TrimSpace(response)
	if response != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			response, expected)
	}
}

func TestUsersCreate(t *testing.T) {
	jsonParams := bytes.NewBuffer([]byte(`{"name":"John"}`))

	req, err := http.NewRequest("POST", "/users", jsonParams)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	handler := ContentType(http.HandlerFunc(UsersCreate))

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got '%v' want '%v'",
			status, http.StatusOK)
	}

	if ctype := responseRecorder.Header().Get("Content-Type"); ctype != jsonHeader {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, jsonHeader)
	}

	expected := `{"name":"John"}`
	response := responseRecorder.Body.String()
	response = strings.TrimSpace(response)
	if response != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			response, expected)
	}

	// Ensure user was added to repo list
	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder = httptest.NewRecorder()
	handler = http.HandlerFunc(UsersIndex)

	handler.ServeHTTP(responseRecorder, req)

	expected = `[{"name":"John"}]`
	response = responseRecorder.Body.String()
	response = strings.TrimSpace(response)
	if response != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			response, expected)
	}
}

func TestUsersCreateWithWrongPayload(t *testing.T) {
	jsonParams := bytes.NewBuffer([]byte(":"))

	req, err := http.NewRequest("POST", "/users", jsonParams)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()
	handler := ContentType(http.HandlerFunc(UsersCreate))

	handler.ServeHTTP(responseRecorder, req)

	if ctype := responseRecorder.Header().Get("Content-Type"); ctype != jsonHeader {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, jsonHeader)
	}

	response := responseRecorder.Body.String()
	response = strings.TrimSpace(response)
	expected := `{"code":400,"text":"payload has wrong format"}`
	if response != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			response, expected)
	}

	status := responseRecorder.Code
	if status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got '%v' want '%v'",
			status, http.StatusOK)
	}
}
