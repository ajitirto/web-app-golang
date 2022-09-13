package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/assests")))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
