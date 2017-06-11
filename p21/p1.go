package main

import "html/template"
import "os"

//::::::: USO DE TEMPLATES :::::::::

const HERO = `
hero name: {{.Name}}
{{range .Emails}}
Emails: {{.}}
{{end}}
{{with .Friends}}
{{range .}}
Friends Name: {{.Name}}
{{end}}
{{end}}
`

//solo el . es uno de todos los items del arreglo

type Friend struct {
	Name string
}

type Hero struct {
	Name    string
	Emails  []string
	Friends []Friend
}

func main() {

	f1 := Friend{"Thor"}
	f2 := Friend{"Hulk"}

	t := template.New("Hero")
	e, err := t.Parse(HERO)
	if err != nil {
		panic(err)
	}

	hero := Hero{
		Name:    "Ironman",
		Emails:  []string{"iron@gmail.com", "heroe@gmail.com"},
		Friends: []Friend{f1, f2},
	}

	err = e.Execute(os.Stdout, hero)
	if err != nil {
		panic(err)
	}
}
