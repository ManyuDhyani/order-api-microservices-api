package main

import (
	"fmt"
	"net/http"
)

// main is the entry point of the program
func main() {
	// Create a new HTTP server
	server := &http.Server{
		// The address to listen on, in this case port 3000 on all network interfaces
		Addr: ":3000",
		// The handler for incoming requests, here we use the basicHandler function
		Handler: http.HandlerFunc(basicHandler),
	}

	// Start the server and listen for incoming HTTP requests
	err := server.ListenAndServe()
	if err != nil {
		// If there is an error starting the server, print it to the console
		fmt.Println("failed to listen to server", err)
	}
}

// basicHandler handles incoming HTTP requests
func basicHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	w.Write([]byte("Server is Live"))
}
