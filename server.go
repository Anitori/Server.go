package main

import (
	"net/http"
	// "fmt"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	//En los strings de a continuación, puedo indicar exactamente la continuación del Dominio para ingresar

	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":8080", nil)
	// handler()
}

// func handler(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }