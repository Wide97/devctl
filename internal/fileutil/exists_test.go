package fileutil

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExists_OK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	exists, err := Exists(context.Background(), ts.URL, "/file.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !exists {
		t.Fatal("expected file to exist")
	}
}

func TestExists_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer ts.Close()

	exists, err := Exists(context.Background(), ts.URL, "/file.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exists {
		t.Fatal("expected file to not exist")
	}
}
