package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "127.0.0.1"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	n := 0

	// Initialize and start the HTTP server on port 8889
	fmt.Println("Server is listening on port 8889")

	// Define a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you are th-%v", n)
		sugar.Infow("request to /",
			"number", n,
		)
		n++
	})

	if err := http.ListenAndServe(":8889", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached %s!", r.URL.Path)
	})

}
