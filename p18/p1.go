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
const tp = `
{{range .}}
Nombre: {{.Nombre}} Edad: {{.Edad}}
{{end}}` //Persona.Nombre, no se  pone le dato Persona porque ya esta en el contexto de la pagina
// el . del range es el parametro que le pasamos

func main() {
	//text template

	persona := []Persona{{"Alejandro", 10},
		{"Maria", 20},
		{"Pedro", 25},
		{"José", 30},
		{"Jaime", 0},
	}

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
