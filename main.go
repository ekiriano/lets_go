package main

import (
	"log"
	"net/http"
)

// Define a home handler function which weites a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the hand
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function
func showSnippet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a specific snippet ..."))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Create a new snippet ..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
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

// Go Serve mux
// Go’s servemux supports two different types of URL patterns: fixed paths
// and subtree paths. Fixed paths don’t end with a trailing slash, whereas
// subtree paths do end with a trailing slash.
// http default servemux is global hence accessible by third party packages which could have
// malicious intent
