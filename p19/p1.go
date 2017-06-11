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
	{{if .Edad}}
		Nombre: {{.Nombre}} Edad: {{.Edad}}
	{{else}}
		Nombre: {{.Nombre}} Edad: {{.Edad}} - Incorrecto
	{{end}}
{{end}}`

//Persona.Nombre, no se  pone le dato Persona porque ya esta en el contexto de la pagina
// el . del range es el parametro que le pasamos
//se puede usar {{else if .Nombre}}
//se imprimen los valores que tengan ead por el if en el template
// el valor 0 en Edad equivale a que no es un número por eso no se imprime jaime

func main() {
	//text template

	persona := []Persona{{"Alejandro", 10},
		{"Maria", 20},
		{"Pedro", 25},
		{"José", 30},
		{"jaime", 0},
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
