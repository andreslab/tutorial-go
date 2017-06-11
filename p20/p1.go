package main

//::::::: USO DE TEMPLATES :::::::::

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
	Pais   string
}

func main() {
	//text template

	persona := Persona{"Alejandro", 10, "Ecuador"}

	t := template.New("persona")
	//e, err := t.ParseFiles("residente","visitante")
	e, err := t.ParseGlob("templates/*.txt") //parsea todo los archivos de la ruta
	if err != nil {
		panic(err)
	}

	err = e.ExecuteTemplate(os.Stdout, "residentes", persona) //SALIDA DEL SISTEMA OPERATIVO, LA CONSOLA
	//especificamos el template que vamos a mostrar, es el nombre del {{define }} que esta en el templates
	if err != nil {
		panic(err)
	}
}
