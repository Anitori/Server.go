package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


func main() {

	fs := http.FileServer(http.Dir("./static"))
	direction := ":8080"

	//En los strings de a continuación, puedo indicar exactamente la continuación del Dominio para ingresar

	// En esta opción mediante un handler, podría escribir algo desde esa misma función

	http.Handle("/", http.StripPrefix("/", fs)) 

	http.HandleFunc("/asd", handler) //Tiene que ir antes de ListenAndServe - TEST
	http.HandleFunc("/qwe", index) //Tiene que ir antes de ListenAndServe - TEST
	http.HandleFunc("/union", unionConHTML)

	/////////////

	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, "Host: ", req.Host)
			fmt.Fprintln(w, "URI: ", req.RequestURI)
			fmt.Fprintln(w, "Method: ", req.Method)
			fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)
			fmt.Fprintln(w, "URL.Query(): ", req.URL.Query())
			fmt.Fprintln(w, "Response: ", req.Response)
			fmt.Fprintln(w, "URL.RawPath: ", req.URL.RawPath)
			fmt.Fprintln(w, "URL.RawQuery: ", req.URL.RawQuery)
			fmt.Fprintln(w, "URL.RawFragment: ", req.URL.RawFragment)
	} )

	log.Fatal(http.ListenAndServe(direction, nil))
}

//////////////////////////////

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

//////////////////////////////

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/test.html")
	if err != nil {
		fmt.Fprint(w, "Página no encontrada!")
	} else {
		template.Execute(w, struct{ Saludo string} {"Vengo de GOLANhG"})
	}
}

/////////////////////////////

func unionConHTML(w http.ResponseWriter, r *http.Request) {
	template:= template.Must(template.ParseFiles("templates/index.html"))
	template.Execute(w, struct{ Saludo string} {"Vengo de GOLANG"})

	}





