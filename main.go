package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached %s!", r.URL.Path)
	})

	// Initialize and start the HTTP server on port 8889
	fmt.Println("Server is listening on port 8889")
	if err := http.ListenAndServe(":8889", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached %s!", r.URL.Path)
	})

	// 访问 /hello 会输出 "Hello, you've reached /hello!"

}
