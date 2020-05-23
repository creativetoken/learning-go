package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to my http server")
	fmt.Fprintln(w, r.UserAgent())

}

func main() {

	var d hotdog

	http.ListenAndServe(":8080", d)

}
