package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello" to the response
	_, err := fmt.Fprintf(w, "Hello")
	if err != nil {
		return
	}
}

func main() {
	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable is not set")
		os.Exit(1)
	}

	// Register the helloHandler for the root URL path
	http.HandleFunc("/", helloHandler)

	// Start the server on the specified port
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
