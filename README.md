# Go Bonobo Http

A Go-Lang HTTP toolbox.

**DISCLAIMER**: This code is part of a formation. Take care!

## Features

### The HttpController

An `HttpController`function is similar to the `HandlerFunc` type from `http` package, 
but returns an error in order to allow error handling in a higher layer.

```go
package MyApi

import "net/http"

func SomethingController(responseWriter http.ResponseWriter, request *http.Request) error {
    result, err := DoSomething()
	if  err != nil{
		return err	
    }

	responseWriter.WriteHeader(200)
	return nil
}

func DoSomething() (string, error) {
	// Do Whatever you want
	// and in case of error, return some error
}
``` 

### Basic error handling
If you work with an `HttpController`, you can use the `HandleError` function to control errors and return a proper 
Internal Server Error. ```HandleError``` implements `HandlerFunc`. That's the point in which we adapt our `HttpController` 
functions to `HandleFunc` so you can use them with your favorite http packages. This is an example of routing set up 
with [`gorilla/mux`](https://github.com/gorilla/mux) for our previous `SomethingController`. 

```go
// ...
router := mux.NewRouter()
var server = http.Server{Addr: ":3000", Handler: router}

router.HandleFunc("/some-path/", SomethingController.HandleError).Methods("GET")

server.ListenAndServe();
// ...

```

If your controller returns a non-nil error value a 500 HTTP Response will be written.  

### Basic Auth

Bonobo also provides a simple way of applying BasicAuth to a controller without modifying it to add 
the credentials check. 

You can use the `ComposeHandler` function to add handlers -of `HandlerFunc` type- to 

```go 
validCredentials := Web.BasicAuthCredentials{"username", "password"}

somethingHandler = Web.ComposeHandler(SomethingController.HandleError, Web.BasicAuth(validCredentials))

router.HandleFunc("/some-path/", somethingHandler).Methods("GET")
```

### Request/Response Logging

Sorry man, this documentation is a work in progress.




