package main

import (
	"fmt"
	"net/http"
)

func main() {

	//servermux propio
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>hola mundo<h1>")
	})

	mux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>hola prueba<h1>")
	})

	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>hola usuario<h1>")
	})

	//crear servidor web
	http.ListenAndServe(":8080", mux) //con nill usa un servemux inetrno

}

//----v2---
//servemux :enrutador de peticiones http
