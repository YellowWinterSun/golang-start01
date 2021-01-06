package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}

	t := template.New("fieldname example")
	//because f1 and f2 is struct(object)
	t, _ = t.Parse(`hello {{.UserName}}!
{{range .Emails}}
    an email {{.}}
{{end}}
{{with .Friends}}
{{range .}}
    my friend name is {{.Fname}}
{{end}}
{{end}}
`)
	t.Execute(os.Stdout, p)
}