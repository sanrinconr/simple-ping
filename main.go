package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		fmt.Print(".")
	})

	fmt.Println("Running!")
	if err := http.ListenAndServe(":8080", s); err != nil {
		fmt.Println(err)
	}
}
