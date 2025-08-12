package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Otro    string `json:"otro"`
	Numero  int    `json:"numero,omitempty"`
	Boolean bool   `json:"boolean,omitempty"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var myString string = "Hello, World!"
	//fmt.Println(myString)
	fmt.Println("Received request for /hello")

	const myInt int = 52

	if myInt == 42 {
		fmt.Println("The answer to life, the universe, and everything is 42")
		myString = "The answer to life, the universe, and everything is 42"
	} else if myInt < 42 {
		fmt.Println("The answer is less than 42")
		myString = "The answer is less than 42"
	} else {
		fmt.Println("The answer is not 42")
		myString = "The answer is not 42"
	}

	response := Message{Message: "Hello from Go API in DevContainer!", Otro: myString}
	fmt.Printf("Response: %+v\n", response)
	// Simulate a delay to mimic processing time
	// time.Sleep(2 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Hola desde Go API en DevContainer!", Otro: myString, Numero: myInt, Boolean: true})
	//json.NewEncoder(w).Encode(response)
}

func holaMundo(w http.ResponseWriter, r *http.Request) {
	var myString string = "Hola, Mundo!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Message: "Hola Mundo", Otro: myString})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", holaMundo)
	http.ListenAndServe(":8080", nil)
}
