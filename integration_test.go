// +build integration

package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestIntegrationGreet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Greet("Jane")))
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("failed to make a request to the server: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", res.StatusCode, http.StatusOK)
	}

	expected := "Hello, Jane!"
	body := make([]byte, len(expected))
	_, err = res.Body.Read(body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if string(body) != expected {
		t.Errorf("unexpected response body: got %s, want %s", string(body), expected)
	}
}
