package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag with name `addr`, a default value of ":4000"
	// and some short help text explaining what the flag controls. 
	// The value of the flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line-flag.
	// This reads in the command-line flag value and assigns it to the addr valiable.
	// You need to call this *before* you use the addr variable otherwise it will always
	// contain the default value of ":4000". If any errors are encountered during parsing
	// the application will be terminated.
	flag.Parse()

	// Use the http.NewServerMux() funtion to initialize a new servermux.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir dunction is relative to the project
	// directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handle for
	// all URL paths that start with "/static/". 
	// For matching paths, we strip the "/static" prefix before the request reaches file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the home function as the handler of the "/" URL pattern.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass i two parameters: the TCP network address to listen on (in this case ":400")
	// and the servermux we just created.
	// If http.ListenAdnServe() returns as error, we use the log.Fatal() function to log error and exit.
	// Note: tht any error returned by http.ListenAndServe() is always non-nil
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}