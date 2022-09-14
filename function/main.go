package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said \"%s\" ", from, message)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var person = Superhero{
			Name:    "Aji Tirto Prayogo",
			Alias:   "Spiderman",
			Friends: []string{"Thor", "Hulk", "Iron Man"},
		}
		var tmpl = template.Must(template.ParseFiles("view.html"))
		
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
