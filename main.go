package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func main() {
	errMsg := color.New(color.FgRed).SprintFunc()
	fs := http.FileServer(http.Dir("./src"))
	heartFs := http.FileServer(http.Dir("./src/heart"))

	// Terminal hyperlink support
	url := "http://localhost:8080/"
	fmt.Printf("Server %s at %s\n", color.GreenString("starting"), color.BlueString("\033]8;;%s\033\\%s\033]8;;\033\\", url, url))

	http.Handle("/", http.StripPrefix("/", fs))           // Serve the main file server at the root
	http.Handle("/heart/", http.StripPrefix("/heart", heartFs)) // Serve the heart file server at /heart/

	err := http.ListenAndServe(":8080", nil) // Use nil to default to DefaultServeMux
	if err != nil {
		log.Fatal(errMsg("ListenAndServe: "), err)
	}
}
