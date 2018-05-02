package main

import (
   "os"
   "text/template"
)

var funcMap = template.FuncMap{
        "Size": Size,
}

type Inventory struct {
	Material string
	Count    uint
}

func (inv Inventory) Color() string {
   return "blue"
}

func main() {
   sweaters := Inventory{"wool", 17}
   tmpl, err := template.New("test").Funcs(funcMap).Parse("{{.Count}} {{Size}} {{.Color}} items are made of {{.Material}}\n")
   if err != nil { panic(err) }
   err = tmpl.Execute(os.Stdout, sweaters)
   if err != nil { panic(err) }
}

func Size() string {
   return "x-large"
}
