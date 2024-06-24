package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of flag
	// will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	// Importantly, we use the flag.

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
