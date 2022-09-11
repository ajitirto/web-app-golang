package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Selamat datang @ajitirto"
	w.Write([]byte(message))

}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = " Halo, selamat datang"
	w.Write([]byte(message))
}


func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started of %s\n", address)


	// menggunakan server 
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())

	}

}