package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SimpleHandler)
	log.Print("Starting the http server...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("Failed to started the http server. Error: ", err)
	}
}

func SimpleHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	type Person struct {
		ID int
		FirstName string
		LastName string
	}

	person := Person{
		ID: 1,
		FirstName: "Antonel",
		LastName: "Pazargic",
	}

	b, _ := json.Marshal(person);

	w.Write(b)
}
