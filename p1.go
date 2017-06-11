package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>hola mundo<h1>")
	})

	//crear servidor web
	http.ListenAndServe(":8080", nil)
}
