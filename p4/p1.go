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

type mensaje struct {
	msg string
}

//Handle propio, sobreescribe el metodo ServeHTTP que lo llama el ServeMux
func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.msg)
}

func main() {

	msg := mensaje{
		msg: "hola mundo de nuevo",
	}

	//servermux propio
	mux := http.NewServeMux()

	//el handleFunc recibe una función que tenga los parametros de un handle
	mux.HandleFunc("/", holaMundo)

	mux.HandleFunc("/prueba", holaPrueba)

	mux.HandleFunc("/usuario", holaUsuario)

	mux.Handle("/hola", msg)

	//crear servidor web
	http.ListenAndServe(":8080", mux) //con nill usa un servemux inetrno

}

//----v2---
//servemux :enrutador de peticiones http
//---v3---
//cualquier objeto puede ser un manejador si implementa Handle
//---v4---
//creamos un handle propio
