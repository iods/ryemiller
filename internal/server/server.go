package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Rye Miller's Website and API's built on the Google App Engine.")
}

func ServerInit() {
	http.HandleFunc("/", Index)

	port := os.Getenv("PORT")
	if port != "8080" {
		port = "8080"
		log.Printf("Defaulting to port %s.\n", port)
	}

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}