package main
import (
	"net/http"
	"fmt"
	"html/template"

)

type Info struct {
	Affiliation string
	Address string
}
type Person struct {
	Name string
	Gender string
	Hobbies []string
	Info Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 20 divisions"
}


func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var person = Person {
			Name : "Aji Tirto Prayogo",
			Gender : "laki-laki",
			Hobbies : []string{"Berenang","badminton", "Nongkrong"},
			Info : Info{"Refactory IS","Jelobo, Wonosari, Klaten"}, 
		}
		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
	

}