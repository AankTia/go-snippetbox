package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to holed the application-wide dependencies for the web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	// Use log.New() to create a logger for writing information messages.
	// This takse three parameters: the destination to write the logs to (os.Stdout),
	// a string preffix for mesage (INFO followed by a tab), and flags to indicate
	// what additional information to include (local date and time).
	// Note that te flag are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application struct, containg the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a http.Server struct.
	// We set the Addr and Handler fields so that the server uses the same network address
	// and routes, and set the ErrorLog field so that the server errorLog logger 
	// in the event of any problems
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Get the servermux containing our routes
	}

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass i two parameters: the TCP network address to listen on (in this case ":400")
	// and the servermux we just created.
	// If http.ListenAdnServe() returns as error, we use the log.Fatal() function to log error and exit.
	// Note: tht any error returned by http.ListenAndServe() is always non-nil
	infoLog.Printf("Starting server on %s", *addr)
	// call teh ListenAndServe() method on our http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
