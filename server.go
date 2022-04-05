package main

import (
	"net/http"
	"log"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	direction := ":8080"

	//En los strings de a continuación, puedo indicar exactamente la continuación del Dominio para ingresar

	http.Handle("/", http.StripPrefix("/", fs)) // En esta opción mediante un handler, podría escribir algo desde esa misma función

	log.Fatal(http.ListenAndServe(direction, nil))
	// handler()
}

// func handler(w http.ResponseWriter, _ *http.Request) {
// 	fmt.Fprintf(w, "TEST TEST TEST TEST")
// }