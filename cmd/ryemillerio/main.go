package main

import (
	"io"
	"net/http"
)

/*
indexHandler Default handler for bootstrapping the application and index page. */
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Rye Miller's Website.")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("localhost:8150", nil)
}