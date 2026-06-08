package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		payload, err := json.Marshal(response{
			Name:        "Hello",
			Description: "World",
		})
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(payload)
	})

	log.Fatal(http.ListenAndServe(":4444", mux))
}
