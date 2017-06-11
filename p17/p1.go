package main

//::::::: USO DE TEMPLATES :::::::::

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
}

//template
const tp = `Nombre: {{.Nombre}} Edad: {{.Edad}}` //Persona.Nombre, no se  pone le dato Persona porque ya esta en el contexto de la pagina

func main() {
	//text template

	persona := Persona{"Alejandro", 30}
	t := template.New("persona")
	e, err := t.Parse(tp) //parseamos para enviar a data
	if err != nil {
		panic(err)
	}

	err = e.Execute(os.Stdout, persona) //SALIDA DEL SISTEMA OPERATIVO, LA CONSOLA
	if err != nil {
		panic(err)
	}
}
