package web

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rye Miller's Website and API's.")
}

func ServerInit() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8150", nil))
}