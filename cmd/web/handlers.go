package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request)  {
	// Exact the value of the id parameter from the query string and try to
	// convert it to an integer using strconv.Atoi() function.
	// If it can't be converted to an integer, or the value is less tahan 1
	// we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter
	fmt.Fprint(w, "Display a specific snippet with ID %d...", id)
}



func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// Use the Header().Set() method to add an `Allow: POST` header to
		// the response header map. The first parameter is the header name, and
		// the second parameter is the header value
		w.Header().Set("Allow", "POST")

		// Use the http.Error() function to send a 405 status code and
		// "Method Not Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("Create a new snippet..."))
}