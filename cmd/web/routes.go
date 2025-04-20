package main

import "net/http"

// The routes() method returns a servermux containing our application routes.
func (app *application) routes() *http.ServeMux{
	// Use the http.NewServerMux() funtion to initialize a new servermux.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir dunction is relative to the project directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handle for
	// all URL paths that start with "/static/".
	// For matching paths, we strip the "/static" prefix before the request reaches file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the home function as the handler of the "/" URL pattern.
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}