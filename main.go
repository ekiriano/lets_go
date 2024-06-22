package main

import (
	"log"
	"net/http"
)

// Define a home handler function which weites a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Go’s servemux
// treats the URL pattern "/" like a catch-all. So at the moment all HTTP
// requests will be handled by the home function regardless of their URL
// path. For instance, you can visit a different URL path like
// http://localhost:4000/foo

// The log.Fatal() function will also call os.Exit(1) after writing the
// message, causing the application to immediately exit.
