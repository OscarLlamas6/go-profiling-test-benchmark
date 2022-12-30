package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website :D!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":8081", nil)
}
