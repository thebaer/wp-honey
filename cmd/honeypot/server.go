package main

import (
	"flag"
	"fmt"
	honey "github.com/writeas/wp-honey"
	"log"
	"net/http"
)

func main() {
	// Get configuration
	portPtr := flag.Int("p", 8080, "Server listening port.")
	flag.Parse()

	// Set up
	err := honey.InitTemplates()
	if err != nil {
		log.Printf("Template init failed: %v", err)
		return
	}

	// Serve application
	http.HandleFunc("/wp-login.php", honey.Handle(honey.NewBee))
	http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
}
