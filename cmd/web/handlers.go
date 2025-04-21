package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/AankTia/go-snippetbox/internal/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
	}

	// Initilize a slice containing the paths to the two files. It's important
	// to mote that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into
	// a template set. If there's an error, we log the detailed error messages
	// and use the http.Error() function to send a generic 500 Internal Server
	// Error response to the user
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execure() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represent any dynamic data we want to pass in, which is now we'll leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Exact the value of the id parameter from the query string and try to
	// convert it to an integer using strconv.Atoi() function.
	// If it can't be converted to an integer, or the value is less tahan 1
	// we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a specific record base on its ID.
	// If no matching record is found, return a 404 Not Found response.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// Use the Header().Set() method to add an `Allow: POST` header to
		// the response header map. The first parameter is the header name, and
		// the second parameter is the header value
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data.
	// We'll remove these later on durring the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\nKobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method,
	// receiving the ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
