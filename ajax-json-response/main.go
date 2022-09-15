package main

import (
	"fmt"
	"net/http"
	"encoding/json"

)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := [] struct {
		Name string
		Age int
		} {
		{ "Richard Grayson", 24 },
		{ "Jason Todd", 23 },
		{ "Tim Drake", 22 },
		{ "Damian Wayne", 21 },
		}
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Header().Set("Content-type", "application/json")
	w.Write(jsonInBytes)


}