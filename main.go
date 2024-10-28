package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		fmt.Print(".")
	})

	port := "8080"
	if customPort := os.Getenv("PORT"); customPort != "" {
		port = customPort
	}

	fmt.Println("Running!")
	if err := http.ListenAndServe(":"+port, s); err != nil {
		fmt.Println(err)
	}
}
