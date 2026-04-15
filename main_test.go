package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	// req := httptest(http.MethodGet, "/root", nil)
	w := httptest.NewRecorder()
	root(w, nil)

	desiredcode := http.StatusOK

	if w.Code != http.StatusOK {
		t.Errorf("bad request expected: %v but returned: %v", desiredcode, w.Code)
	}

	expectedMessage := "this is the home page\n"

	if !bytes.Equal([]byte(expectedMessage), w.Body.Bytes()){
		t.Errorf("bad response expected: %q but got: %q", expectedMessage, w.Body.String())
	}
}

func TestPing(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "/ping?user=jack", nil)

	w := httptest.NewRecorder()

	ping(w, req)

	desiredcode := http.StatusOK

	if desiredcode != w.Code {
		t.Errorf("Bad response expected: %v but received: %v", desiredcode, w.Code)
	}
	expectedMessage := []byte("Hello jack!\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad response body expected: %q but received: %v", expectedMessage, w.Body.String())
	}
}

func TestRegister(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "/register?user=jack", nil)

	desiredcode := http.StatusOK

	w := httptest.NewRecorder()
	register(w, req)
	if w.Code != desiredcode {
		t.Errorf("bad response expected: %v but received: %v body %v", desiredcode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Welcome jack!\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad body expected: %q but receives: %q", expectedMessage, w.Body.String())
	}
}
