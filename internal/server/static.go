package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func serveFiles(directory string) http.Handler {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v\n", err)
	}

	return http.FileServer(http.Dir(fmt.Sprintf("%s/%s/", dir, directory)))
}



func Static(directory string) func(w http.ResponseWriter, r *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		path = filepath.Join(directory, path)

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(directory, "index.html"))
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		serveFiles(directory).ServeHTTP(w, r)
	})
}
