// HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation
// A context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// a context.Context is created for each request by the net/http machinery, and is available with the COntext() method
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	// wait for a few seconds before sending a reply to the client to simulate some work the server is doing
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	// while working, keep an eye on the context's Done() channel for a signal that we should cancel the work and return as soon as possible
	case <-ctx.Done():
		// The context's Err() method returns an error that explains why the Done() channel was closed
		err := ctx.Err()
		fmt.Println("server: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// we register hello handler on the "/hello" route, and start serving
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}


// output
// success case:
// $ go run src/80_context/context.go &                                                                                                                  [23:31:05]
// [1] 872383
 
// $ starting server                                                                                                                                     [23:31:09]

// $ curl localhost:8090/hello                                                                                                                           [23:31:16]
// server: hello handler started
// server: hello handler ended

// error case:
// $ go run src/80_context/context.go &                                                                                                                  [23:33:51]
// [2] 875466
 
// $ curl localhost:8090/hello                                                                                                                           [23:34:04]
// server: hello handler started
// ^Cserver:  context canceled

// server: hello handler ended