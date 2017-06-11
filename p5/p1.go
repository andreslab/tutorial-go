package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

	//con el & indicamos que creamos un puntero de una estructura Server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second, //tiempo de espera de lectura (10 segundos)
		WriteTimeout:   10 * time.Second, //tiempo de espera de escritura
		MaxHeaderBytes: 1 << 20,          //numero en bytes .. 1 * 2, 20 veces, 1 Mb
	}

	log.Println("Listening....")
	log.Fatal(server.ListenAndServe()) //cuando habria un error se detendria el servidor
	//tamb se puede solo:  server.ListenAndServe()
}

//----v2---
//servemux :enrutador de peticiones http
//---v3---
//cualquier objeto puede ser un manejador si implementa Handle
//---v4---
//creamos un handle propio
//----v5----
//crearemos nuestro propio servidor ... ya no le pasamso parametros a ListenAndServe porque se paso en la estructura del server
