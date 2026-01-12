package ping

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheck_OK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	if err := Check(context.Background(), ts.URL); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestCheck_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer ts.Close()

	if err := Check(context.Background(), ts.URL); err == nil {
		t.Fatal("expected error, got nil")
	}
}
