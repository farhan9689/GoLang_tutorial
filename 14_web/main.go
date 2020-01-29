package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello world</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello again, you are in about page</h1>")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("server is starting...")
	http.ListenAndServe(":3000", nil)
}
