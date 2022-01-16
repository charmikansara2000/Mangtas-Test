package main

import (
	"log"
	"net/http"

	"mangtas/controllers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/frequency", controllers.Execute).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))

}
