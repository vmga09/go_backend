package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Hola desde Go API en DevContainer!"})
}

func holaMundo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Hola Mundo"})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", holaMundo)
	http.ListenAndServe(":8080", nil)
}
