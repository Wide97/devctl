package mockserver

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/sys/info", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{
			"os":      runtime.GOOS,
			"arch":    runtime.GOARCH,
			"runtime": runtime.Version(),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(info); err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/file/exists", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("path")
		if path == "" {
			http.Error(w, "missing path", http.StatusBadRequest)
			return
		}
		if _, err := os.Stat(path); err == nil {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})

	return mux
}
