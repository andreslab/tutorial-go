package main

import (
	"net/http"
	"time"

	"log"

	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"` // <- anotacion, las anotaciones son para un valor si es nullo el parametro
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

var noteStore = make(map[string]Note) //de tipo Note

var id int

//GetNoteHandler - Get - /api/notes
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, value := range noteStore {
		notes = append(notes, value)
	}
	//cabecera http
	w.Header().Set("Content-Type", "application/json") //llamar antes del WriteHeader
	j, err := json.Marshal(notes)                      //convertir estructura de go en json
	if err != nil {
		panic(err) // esto detiene el servidor .. usar solo en debug
	}
	//w.WriteHeader(200) // 200 code OK
	w.WriteHeader(http.StatusOK) //statusOK significa code 200

	//configurar el cuerpo
	w.Write(j)
}

//PostNoteHandler - Post - /api/notes
func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note) //devuelve decodificador el json de la peticion del usuario y lo rellene en la estructura note
	if err != nil {
		panic(err)
	}
	note.CreatedAt = time.Now()
	id++                  //incrementa el id
	k := strconv.Itoa(id) //pasar entero a string porque el key de noteStore es string
	noteStore[k] = note

	//:::: devolver json de respuesta con la fecha de creacion:::::::::

	//cabecera http
	w.Header().Set("Content-Type", "application/json") //llamar antes del WriteHeader
	j, err := json.Marshal(note)                       //convertir estructura de go en json
	if err != nil {
		panic(err) // esto detiene el servidor .. usar solo en debug
	}
	//w.WriteHeader(200) // 200 code OK
	w.WriteHeader(http.StatusCreated) //statusOK significa code 201

	//configurar el cuerpo
	w.Write(j)

}

//PutNoteHandler - Put - /api/notes
func PutNotesHandler(w http.ResponseWriter, r *http.Request) {

	//paquete vars que extrae variables o parametros
	vars := mux.Vars(r)
	k := vars["id"] //convirtio en un slide de string y le pasamos el key para extraer parámetro
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	//ya sabemos el id de la nota y los datos

	if note, ok := noteStore[k]; ok {
		//la funcion ok = sirve apra ver si el dato existe y se ejecuta antes del if
		//el ok funcinona como un bool
		noteUpdate.CreatedAt = note.CreatedAt //usamos la misma fecha cuando se creo, no de la actualziacion
		delete(noteStore, k)                  //le pasamos el inidice k que vamos a borrar de noteStore
		noteStore[k] = noteUpdate

	} else {
		log.Printf("no encontramos el id %s", k)
	}

	w.WriteHeader(http.StatusNoContent) // code 204

}

//DeleteNoteHandler - Get - /api/notes
func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {

	//paquete vars que extrae variables o parametros
	vars := mux.Vars(r)
	k := vars["id"] //convirtio en un slide de string y le pasamos el key para extraer parámetro
	//ya sabemos el id de la nota y los datos

	if _, ok := noteStore[k]; ok { //deseche la variable note de la funcion anterior porque no se la usa

		delete(noteStore, k) //le pasamos el inidice k que vamos a borrar de noteStore
		//noteStore[k] = noteUpdate

	} else {
		log.Printf("no encontramos el id %s", k)
	}

	w.WriteHeader(http.StatusNoContent) // code 204
}

func main() {
	r := mux.NewRouter().StrictSlash(false) //no las considera iguales a las rutas
	r.HandleFunc("/api/notes", GetNotesHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNotesHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNotesHandler).Methods("PUT") //{id} variable
	r.HandleFunc("/api/notes/{id}", DeleteNotesHandler).Methods("DELETE")

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Printf("LISTENING")
	server.ListenAndServe()
}
