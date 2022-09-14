package main
import (
	"fmt"
	"net/http"
	"html/template"

)

const view string = `<html>
	<head>
		<title> Template</title>
	</head>
	<body>
		<h1>hello</h1>
	</body>
	</html>
	`

func main () {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w,nil);err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000",nil)
	
}