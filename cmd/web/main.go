package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of flag
	// will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Importantly, we use the flag.

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
