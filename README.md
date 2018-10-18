# Go Bonobo Http

A Go-Lang HTTP toolbox.

**DISCLAIMER**: This code is part of a formation. Take care!

## Features

### The HttpController

`HttpController` defines an interface for controllers in a similar way to the
`http.handleFunc` but also adding the possibility to return a error. 

```go
package MyApi

import "net/http"

type SomethingController struct { }

func (controller SomethingController) Action(responseWriter http.ResponseWriter, request *http.Request) error {
    result, err := DoSomething()
    if  err != nil {
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
If you work with an `HttpController`, you can use the `HandleErrorMiddleWare` to control errors and return a proper 
Internal Server Error. `HandleErrorMiddleWare` implements `http.handlerFunc`. That's the point in which we adapt our
`HttpController.Action` functions to `HandleFunc` so you can use them with your favorite http packages. This is an example of routing set up 
with [`gorilla/mux`](https://github.com/gorilla/mux) for our previous `SomethingController`. 

```go
// ...
router := mux.NewRouter()
var server = http.Server{Addr: ":3000", Handler: router}
errorHandler := Web.HandleErrorMiddleWare{SomethingController{}}
router.HandleFunc("/some-path/", errorHandler.Action).Methods("GET")

server.ListenAndServe();
// ...

```

If your controller returns a non-nil error value a 500 HTTP Response will be written.  

### Basic Auth

Bonobo also provides a simple way of applying BasicAuth to a controller without modifying it to add 
the credentials check. 

You can use the `ComposeHandler` function to add handlers -of `HandlerFunc` type- to other handler functions. 

```go 
validCredentials := Web.BasicAuthCredentials{"username", "password"}

composedHandler = Web.ComposeHandler(errorHandler.Action, Web.BasicAuth(validCredentials))

router.HandleFunc("/some-path/", composedHandler).Methods("GET")
```

### Request/Response Logging

Sorry man, this documentation is a work in progress.




