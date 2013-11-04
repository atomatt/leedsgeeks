package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port = flag.String("port", "5000", "http port")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", index)
	http.ListenAndServe(":"+*port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}
