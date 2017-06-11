package main

import (
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Mensaje desde e metodo GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Mensaje desde e metodo POST")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Mensaje desde e metodo DELETE")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Mensaje desde e metodo PUT")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	/*
		rutas iguales con false
		/api/user
		/api/user/
	*/

	// el crud (create, read, update, delete) para la api
	r.HandleFunc("/api/users", GetUsers).Methods("GET")   //obtiene un usuario
	r.HandleFunc("/api/users", PostUsers).Methods("POST") //crea un usuario
	r.HandleFunc("/api/users", PutUsers).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("LISTENING")
	server.ListenAndServe()
}

//----v2---
//servemux :enrutador de peticiones http
//---v3---
//cualquier objeto puede ser un manejador si implementa Handle
//---v4---
//creamos un handle propio
//----v5----
//crearemos nuestro propio servidor ... ya no le pasamso parametros a ListenAndServe porque se paso en la estructura del server
//---v6----
//handle fileServe, para servir archivos html, js, css
//con go run se crea un ejecutable temporal en otro directorio por lo que se debe crear ejecutable y luego ejecutarlo
//----v7----
/*
diferencia entre

//funciones del metodo http
http.Handle() //metodo que indica que utilice un manejador
http.HandleFunc()

//interfaz
http.Handler //la interfaz Handle, con funciones que hay q implementar
http.HandlerFunc // la interfaz, es un type que recibe ResponseWriter y un request
*/
//-----v8-----
//instalar gorilla mux
