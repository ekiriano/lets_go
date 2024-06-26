package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of flag
	// will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// Importantly, we use the flag.

	// Swap the route declarations to use the application struct's methods as t
	// handler functions.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
