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
