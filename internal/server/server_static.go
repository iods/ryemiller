package server

import (
	"log"
	"net/http"
	"os"
)

func Run() {

	staticpath := "web/dist"

	r := http.NewServeMux()
	r.HandleFunc("/", Static(staticpath))


	port := os.Getenv("PORT")
	if port != "8080" {
		port = "8080"
		log.Printf("Defaulting to port %s.\n", port)
	}

	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatal(err)
	}
}