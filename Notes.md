# Notes

LET'S GO! (A Step-by Step Guide to Creating Fast, Secure, And Maintanable Web Application with Go)

# 1. Introduction

## 1.1. Prerequisites

1. Go

   - (Installing Go on Mac OS X)[https://golang.org/doc/install#tarball]
   - (Installing Go on Windows)[https://golang.org/doc/install#windows]
   - (Installing Go on Linux)[https://golang.org/doc/install#tarball]
   - (Removing an old version of Go)[https://golang.org/doc/manage-install#uninstalling]

2. (curl)[https://curl.haxx.se/]
   Tool for working with HTTP requests and response from teminal.

---

# 2. Foundations

## 2.1. Project setup and creating a module

### Module

- If you’re not already familiar with Go modules, you can think of a module path as basically being a canonical name or identifier for your project.
- You can pick almost any string as your module path, but the important thing to focus on is uniqueness. To avoid potential import conflicts with other people’s projects or the standard library in the future, you want to pick a module path that is globally unique and unlikely to be used by anything else. In the Go community, acommon convention is to base your module paths on a URL that you own.
- To create / init module, make sure that you're in the root of the directory and then run teh `go mod init` command, passing in your module path is a parameter, like do :
  ```bash
  go mod init github.com/AankTia/go-snippetbox
  ```
- Use the `go run .` command in our terminal to compile and execute the code in current directory.

## 2.2. Web application basics

Three absolute essentials:

1. **_handler_**

   If you're comming from an MVC-backgorund, you can think of handlers as being a bit like _controllers_. They're responsible for executing application logic and for writting HTTP response headers and bodies.

2. **_router_**

   _servermux_ in Go Terminology. This stores a mapping between the URL patterns for your application and the corresponding handlers.

   Usually you have one servermux for your application containing all your routes.

3. **_web server_**

   On eof the great things about Go is that you can establish a web server and listen for incoming requests _as part of your application self_. You don't need an external thrd-party server link Nginx or Apache.

### Network addresses

The TCP network address that you pass to `http.ListenAndServe()` should be in format `host:port`. If you omit the host (only the port like `:4000`) then the server will listen on all your computer's available network interface.

Generaly, you only need to specify a host in the address if your computer has multiple network interface and tou want to listen on just one of them.

### Using `go run`

- During development the `go run` command is convienient way to tru out your code. It's essentially a shortcut that compiles your code, creates an executable binary in your `/tmp` directory, and then runs this binary in one step.
- It accepts either a space-separated list of `.go` files, the path to a specific package (where the `.` character represents your current directory), or the full module path. For our application at the moment, the tree following command are all equivalent:

```bash
go run .
go run main.go
go run github.com/AankTia/go-snippetbox
```

## 2.3. Routing requests

### Fixed Path and Subtree Patterns

Go servermux supports two different types of URL patterns:

- **_fixed paths_**

  _Don't end_ with trailing slash, example:

  ```
  /snippet/view
  /snippet/create
  ```

- **_subtree paths_**

  _Do end_ with a traling slash, example:

  ```
  /
  /static/
  ```

  Subtree path pattern are atched (and the corresponding handler calles) whenever the `start` of a request URL path. You can thnk of subtree paths as acting a bit like they have a wildcard at the end, like `/**` or `/static/`

  This help explains why the `/` pattern is acting like a catch-all. The pattern essentially means _match a single slash, followed by anything (or nothing all)_

### DefaultServeMux

If you've been working with Go for a while you migh have come accross the `http.Handle()` and `http.HandleFunc()` functions. Theses allow you to register routes `without declaring a servermux`.

Behind the scenes, tehese functions register their routes with something called the _DefaultServeMux_. Initialized by default and stored in `net/http` global variable.

```go
var DefaultServeMux = NewServeMux()
```

Although this approach can make your code slightly shorter, **_don't use it for production applications_**. Because `DefaultServeMux` is a global variable, any package can access it and register a route, including any third-party packages that your application imports. Id one of those third-party packages is compromised, they could use `DefaultServeMux` to expose a malicious handler of the web.

So, for the sake of sevurity, it's generally a good idea to avoid `DefaultServeMux` and the corresponding helper functions.

## 2.4. Customizing HTTP headers

### `http.Error` Shortcut

If you want to send a non-`200` status code and a plain text response body, then it's a good opportunity to use `http.Error` schortcut. This is a lighweight helper function which takes a given message and status code, then call the `w.WriteHeader()` and `w.Write` methods begin-the-scene for us.

### System-generated headers and content sniffing

When sending a response Go will automatically set three _system_generated_:

1. `Date`
2. `Content-Lenght`
3. `Content-Type`
   ...

### Manipulating The Header Map

We use `w.Header().Set()` to ass a new header to the response header map. But there's also `Add()`, `Del()`, `Get()` and `Values()` methods that can to read and manipulate the header map.

```go
// Set a new cache-control header. If an existing "Cache-Control" header exists
// it will be overwritten.
w.Header().Set("Cache-Control","public, max-age=31536000")

// In contrast, the Add() method appends a new "Cache-Control" header and can
// be called multiple times.
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")

// Delete all values for the "Cache-Control" header.
w.Header().Del("Cache-Control")

// Retrieve the first value for the "Cache-Control" header.
w.Header().Get("Cache-Control")

// Retrieve a slice of all values for the "Cache-Control" header.
w.Header().Values("Cache-Control")
```

### Header canonicalization

The header name will be always bes canonicalized using the `textproto.CanonicalMIMEHeaderKey()` function.

This converts the first letter and any letter following a hypen to upper case, and the rest of the letters to lowercase.This has the practical implementation that when calling these methods the header name is _sace-insensitive_

If you nees to avoid this cannonicalization behaviour, you can edit the header name is `case-insensitive`.

## 2.5. URL query strings

To retrive the value of the `id` parameter from URL query string, whish we can do using `r.URL.Query().Get()` method. This will always return a string value of a parameter, or the empty string `""` if no matching parameter exists.

## 2.6. Project structure and organization

The structure of project repository should look like this:

```
/project-name
    /cmd
        /web
            handlers.go
            main.go
    /internal
    /ui
        /html
        /static
    go.mod
```

- `cmd` directory

  Contain the _application specific_ code for the executable applications in the project.

- `internal` directory

  Contain the ancillary _non-application-specific_ code used in the project. Use it to hold potentially reusable code like validation helpers and the SQL database models for the project.

- `ui` directory

  Contain the `user-interface assets` used by the web application. Specifically, the `ui/html` directory will contain HTML, and the `ui/static` directory will contain static files (like css and images)

**Benefits using this structure**

1. It gives a clean speration between Go and non-Go assets. This can make things easier to manage when it comes to building and deploying your application in the future

2. It scales really nicely if you want to add another executable application to your project. For example, you might want to add a CLI (Command Line Interface) to automate some administrative tasks in the future. With this sctucture, you could create this CLI application under `cmd/cli` and it will be able to import and reuse all the code you've written under the `internal` directory.

## 2.7. HTML templating and inheritance

### Template composition

To save us typing and prevent duplication, it's a good idea to create a _base_ (or _master_) template which contains this shared content, which we can then _compose_ with the page-specific markup for the individual pages.

We're using the `{{define "base"}}...{{end}}` action to define a distinct `named template` called `base`, which contains the content we want to appear on ever page.

> **Note** :
>
> If you're wondering, the dot at the end of the `{{template "title" .}}` action represents any dynamic data that you want to pass to the invoked template.

## 2.8. Serving static files

### `http.Fileserver` handler

Go's `net/http` package ships with a build-in `http.FileServer` handler which you can use to serve files over HTTP from a specific directory.

> **_Important Note:_**
>
> Once the application is up-and-running, `http.FileServer` probably wont'be reading file from disk.
> Both Windows and Unix-based operating system cache recently-used files in RAM, so (for frequently-served files at least) it's likely that `http.FileServer` will be serving then from RAM rather than making the relatively slow round-trip to your har disk.

### Serving single files (`http.ServeFile`)

Sometime you might want to serve a single file from within a handler. For this there's the `http.ServeFile()` function, which you can use like so :

```go
func downloadHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./ui/ststic/file.zip")
}
```

> **_Warning_**
>
> `http.ServeFile()` does not automatically sanitize the file path. If you're constructiong a file path form untrusted user input, to avoid directory traversal attacks **_you must_** sanitize the input with _[filepath.Clean()](https://pkg.go.dev/path/filepath/#Clean)_ before using it.

### Disabling directory listing

If you want to disable directory listings there are a few different approaches you can take.

**_The Simple Way:_** Add a blank `index.html` file to the specific directory that you want to dosable listings for. This will then be served instead of the directory listing, and the user will get a `200 OK` response with no body.

If you want to do this for all directories under `.ui/static` you can use the command:

```bash
find ./ui/static -type -d exec touch {}/index.html \;
```

**_A more complicated_** (but arguably better) solution is to create a custom implementation of `http.FileSystem`, and have it return an `os.ErrorNotExist error for any directories`
...

## 2.9. The http.Handler interface

---

# 3. Configuration and error handling

## 3.1. Managing configuration settings

### Command-line flags

In Go, a common and idiomatic way to manage configuration settings is to use `command-line` flags when starting an application.

For example:

```bash
go run ./cmd/web -addr=":80"
```

The easiest wat to accept and parse a command-line flag form your applcation is with a line of conde like this:

```go
addr := flag.String("addr", ":4000", "HTTP network address")
```

> **_Note_**
>
> Ports 0-1023 are restricted and (typically) can only be used by services which have root privileges.
> If you try to use one of these port you should get a `bind: permission denied` error message on start-up

### Automated help

Anoter great feature is that you can use the `-help` flag to list all the available command-line flags for an application and their accompanying help text.

```bash
go run ./cmd/web -help
```

### Environment variables

You can store your configuration settings on environment variables and access them directly from your application by using the `os.Getenv()` function like so:

```go
addr: os.Getenv("SNIPPETBOX_ADDR")
```

### Pre-existing variables

It's possible to parse command-line flag values into the memory addresses of pre-ecisting variables, using the `flag.StringVar()`, `flag.IntVar()`, `flag.BoolVar()` and oher function.

This can be useful if you want to strore all your configuration settings in a single struct. As a rough example:

```go
type config struct {
    addr        string
    staticDir   string
}

...

var cfg config

flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network adress")
flag.StringVar(&cfg.staticDir, "static-dir", ".ui/static", "Path to static asset")

flag.Parse()
```

## 3.2. Leveled logging

The `log.Fatal()` function will also call `os.Exit(1)` after writting teh message, causing the application to immediately exit.

> **_Tip_**
>
> If tou want to include the full file path in your log output, instead of just the file name, you can use the `log.Llongfile` flag instead of `log.Lshortfile` when creating your system logger.
> You can also force your logger to use UTC datetimes (instead of local ones) by adding the `log.LUTC` flag.

### Decoupled logging

A big benefit of loggin your messages to the standard streams (stdout and stderr) are decoupled. Your application itself isn't concered with the routing or storage of the logs. and that can make it easier to manage the logs differenty depending on the environment.

Duting de development, it's easy to view the log output because the standard streams are displayed in the terminal.

In staging or production environments, you can redirect the streams to a final destination fr viewing and acrhival. This destination could be on-disk files, or a loging service.

For example, we could redirect the stdout and stderr streams to on-disk files when staring the application like so:

```bash
go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```

> **_Note**
>
> Using the double arrow `>>` will append to an existing file, instead of truncating it when starting the application

## 3.3. Dependency injection

## 3.4. Centralized error handling

## 3.5. Isolating the application routes

---

# 4. Database-driven responses

## 4.1. Setting up MySQL

## 4.2. Installing a database driver

## 4.3. Modules and reproducible builds

## 4.4. Creating a database connection pool

## 4.5. Designing a database model

## 4.6. Executing SQL statements

## 4.7. Single-record SQL queries

## 4.8. Multiple-record SQL queries

## 4.9. Transactions and other details

---

# 5. Dynamic HTML templates

## 5.1. Displaying dynamic data

## 5.2. Template actions and functions

## 5.3. Caching templates

## 5.4. Catching runtime errors

## 5.5. Common dynamic data

## 5.6. Custom template functions

---

# 6. Middleware

## 6.1. How middleware works

## 6.2. Setting security headers

## 6.3. Request logging

## 6.4. Panic recovery

## 6.5. Composable middleware chains

---

# 7. Advanced routing

## 7.1. Choosing a router

## 7.2. Clean URLs and method-based routing

---

# 8. Processing forms

## 8.1. Setting up a HTML form

## 8.2. Parsing form data

## 8.3. Validating form data

## 8.4. Displaying errors and repopulating fields

## 8.5. Creating validation helpers

## 8.6. Automatic form parsing

---

# 9. Stateful HTTP

## 9.1. Choosing a session manager

## 9.2. Setting up the session manager

## 9.3. Working with session data

---

# 10. Security improvements

## 10.1. Generating a self-signed TLS certificate

## 10.2. Running a HTTPS server

## 10.3. Configuring HTTPS settings

## 10.4. Connection timeouts

---

# 11. User authentication

## 11.1. Routes setup

## 11.2. Creating a users model

## 11.3. User signup and password encryption

## 11.4. User login

## 11.5. User logout

## 11.6. User authorization

## 11.7. CSRF protection

---

# 12. Using request context

## 12.1. How request context works

## 12.2. Request context for authentication/authorization

---

# 13. Optional Go features

## 13.1. Using embedded files

## 13.2. Using generics

---

# 14. Testing

## 14.1. Unit testing and sub-tests

## 14.2. Testing HTTP handlers and middleware

## 14.3. End-to-end testing

## 14.4. Customizing how tests run

## 14.5. Mocking dependencies

## 14.6. Testing HTML forms

## 14.7. Integration testing

## 14.8. Profiling test coverage

---

# 15. Conclusion

---

# 16. Further reading and useful links

---

# 17. Guided exercises

## 17.1. Add an 'About' page to the application

## 17.2. Add a debug mode

## 17.3. Test the snippetCreate handler

## 17.4. Add an 'Account' page to the application

## 17.5. Redirect user appropriately after login

## 17.6. Implement a 'Change Password' feature
