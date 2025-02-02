package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "json")
		w.Write([]byte("Hello World"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Error getting hostname:", err)
			return
		}

		w.Header().Set("Content-Type", "json")
		w.Write([]byte(hostname))
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Server is running on http://localhost:80")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
