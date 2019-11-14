package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// https://golang.org/pkg/net/http/

// Payload ...
type Payload struct {
	Stuff Data
}

// Data ...
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

// Fruits ...
type Fruits map[string]int

// Vegetables ...
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		log.Fatal(err)
	}

	// added newline to end of response
	// removes % character when testing in zsh
	fmt.Fprintf(w, string(response)+"\n")
}

func serveHealth(w http.ResponseWriter, r *http.Request) {
	reponse := "This app is healthy\n"
	fmt.Fprintf(w, reponse)
}

func main() {
	http.HandleFunc("/health", serveHealth)
	http.HandleFunc("/", serveRest)

	http.ListenAndServe("0.0.0.0:1337", nil)
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
