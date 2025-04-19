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

## 2.3. Routing requests

## 2.4. Customizing HTTP headers

## 2.5. URL query strings

## 2.6. Project structure and organization

## 2.7. HTML templating and inheritance

## 2.8. Serving static files

## 2.9. The http.Handler interface

---

# 3. Configuration and error handling

## 3.1. Managing configuration settings

## 3.2. Leveled logging

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
