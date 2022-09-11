package main

import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "welcome"
	w.Write([]byte(message))


}

func handlerHello( w http.ResponseWriter, r *http.Request) {
	var message = "Hello World!"
	w.Write([]byte(message))

}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address  = "localhost:8000"
	fmt.Println("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
		
	}
}

