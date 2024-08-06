package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Define the domains you want to serve
	domains := []string{"yahallo.me", "www.yahallo.me"}

	// Configure the autocert manager
	m := &autocert.Manager{
		Cache:      autocert.DirCache("/var/www/.cache"), // specify a directory to store the certificates
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
	}

	// Create the server with the autocert manager
	srv := &http.Server{
		Addr: ":443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle requests here
			if r.URL.Path == "/heart" {
				http.ServeFile(w, r, "./src/heart")
				return
			}
			http.FileServer(http.Dir("./src")).ServeHTTP(w, r)
		}),
		TLSConfig: m.TLSConfig(),
	}

	// Redirect HTTP to HTTPS
	go func() {
		log.Fatal(http.ListenAndServe(":80", m.HTTPHandler(nil)))
	}()

	log.Println("Server starting on :443")
	err := srv.ListenAndServeTLS("", "") // Using autocert, no need to specify cert paths
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
