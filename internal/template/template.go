package template

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var templates *template.Template

func Compile() {
	var templateFiles []string

	dir := os.Getenv("TEMPLATES_DIR")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("could not read from $TEMPLATES_DIR [%s]: %v", dir, err)
	}
	for _, file := range files {
		templateFiles = append(templateFiles, path.Join(dir, file.Name()))
	}

	templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatalf("could not parse html templates: %v", err)
	}
}