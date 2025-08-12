package main

import (
	"container/list"
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
	/*
		Definicion de un array en Go
		Un array es una colección de elementos del mismo tipo, con un tamaño fijo.
		Los arrays en Go son de tamaño fijo y no se pueden redimensionar.
		Para colecciones dinámicas, se utilizan slices.
	*/
	//Definiendo array de numeros
	var numeros = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	// Definiendo un array de strings
	nombres := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(nombres)
	// Sin especificar el tamaño del array
	var frutas = [...]string{"Manzana", "Banana", "Cereza"}
	fmt.Println(frutas)
	fmt.Println("Array de frutas:", frutas[0])

	frutas[0] = "Pera" // Modificando el primer elemento del array
	fmt.Println("Array de frutas modificado:", frutas)

	var anyArray = [3]any{"Texto", 123, true}
	fmt.Println("Array con tipo any:", anyArray)

	// SLICES
	/*
		Definicion de un slice en Go

		Un slice es una colección dinámica de elementos del mismo tipo.
		Los slices son más flexibles que los arrays, ya que pueden crecer y reducirse en tamaño.
		Un slice es una referencia a un array subyacente, lo que significa que los cambios en el slice afectan al array original.
	*/
	var miSlice = []string{"Rojo", "Verde", "Azul"}
	fmt.Println("Slice original:", miSlice)
	miSlice = append(miSlice, "Amarillo")
	fmt.Println("Slice modificado:", miSlice)

	/*
		Definicion de un map en Go (Diccionario en otros lenguajes)
		Un map es una colección de pares clave-valor, donde cada clave es única.
		Los maps son útiles para almacenar datos que se pueden acceder por una clave.
	*/

	var miMapa = map[string]int{
		"uno": 1,
		"dos": 2}

	fmt.Println("Mapa original:", miMapa)
	miMapa["tres"] = 3 // Agregando un nuevo par clave-valor
	fmt.Println("Mapa modificado:", miMapa)
	// Declarando un mapa con tipo any
	// El tipo any es un alias para interface{}
	// Permite almacenar cualquier tipo de dato en el mapa
	// Es útil cuando no se conoce el tipo de dato de antemano
	// o cuando se quiere permitir una mayor flexibilidad en los tipos de datos.
	var miMapa2 = map[string]any{"nombre": "Juan", "edad": 30, "activo": true}
	fmt.Println("Mapa con tipo any:", miMapa2)
	miMapa2["pi"] = 3.14 // Agregando un nuevo par clave-valor con tipo float64
	fmt.Println("Mapa con tipo any modificado:", miMapa2)

	// LISTA
	myList := list.New()
	myList.PushBack("Elemento 1")
	myList.PushBack("Elemento 2")
	myList.PushBack("Elemento 3")
	fmt.Println(myList.Front().Value) // Imprime el primer elemento de la lista

	var myString string = "Hello, World!"
	//fmt.Println(myString)
	fmt.Println("Received request for /hello")

	const myInt int = 22
	const myFloat float64 = 3.14

	if myInt == 42 || myFloat < 3.2 {
		fmt.Println("The answer to life, the universe, and everything is 42")
		myString = "The answer to life, the universe, and everything is 42 and float is less than 3.2"
	} else if myInt < 42 && myFloat < 3.2 {
		fmt.Println("The answer is less than 42")
		myString = "The answer is less than 42 and float is less than 3.0"
	} else {
		fmt.Println("The answer is not 42")
		myString = "The answer is not 42"
	}

	response := Message{Message: "Hello from Go API in DevContainer!", Otro: myString}
	fmt.Printf("Response: %+v\n", response)
	// Simulate a delay to mimic processing time
	// time.Sleep(2 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(Message{Message: "Hola desde Go API en DevContainer!", Otro: myString, Numero: myInt, Boolean: true})
	json.NewEncoder(w).Encode(miMapa2)
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
