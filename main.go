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

	locationCertificate := os.Getenv("CERT_FILE_LOCATION")
	locationKey := os.Getenv("KEY_FILE_LOCATION")

	fmt.Printf("Running in port %s!\n", port)
	fmt.Println(locationCertificate)
	fmt.Println(locationKey)

	if err := http.ListenAndServeTLS(":"+port, locationCertificate, locationKey, s); err != nil {
		fmt.Println(err)
	}
}
