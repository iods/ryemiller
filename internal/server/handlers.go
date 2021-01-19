package server

import (
	"html/template"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dir, _ := os.Getwd()
	templates := template.Must(template.ParseFiles(dir + "/web/template/index.gohtml"))

	if err := templates.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}