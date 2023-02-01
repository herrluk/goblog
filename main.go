package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "<h1>Hello, 这里是 goblog </h1>")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}
