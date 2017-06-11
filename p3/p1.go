package main

import (
	"fmt"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hola mundo<h1>")
}
func holaPrueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hola mundo<h1>")
}
func holaUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hola mundo<h1>")
}

func main() {

	//servermux propio
	mux := http.NewServeMux()

	//HandleFunc es uno por defecto del serveMux creado
	mux.HandleFunc("/", holaMundo)

	mux.HandleFunc("/prueba", holaPrueba)

	mux.HandleFunc("/usuario", holaUsuario)

	//crear servidor web
	http.ListenAndServe(":8080", mux) //con nill usa un servemux inetrno

}

//----v2---
//servemux :enrutador de peticiones http
//---v3---
//cualquier objeto puede ser un manejador si implementa Handle
//creamos un handle propio
