package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message struct {
			Message string `json:"message"`
		}
		message.Message = "Hello, Docker!"
		json.NewEncoder(w).Encode(message)
	})

	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}

	log.Println("Listening on port 8080")
	s.ListenAndServe()
}
