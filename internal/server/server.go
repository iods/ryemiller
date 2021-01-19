package server

import (
	"log"
	"net/http"
	"os"

	"github.com/iods/ryemiller.io/internal/handler"
	"github.com/julienschmidt/httprouter"
)

func Run() {
	r := httprouter.New()
	r.GET("/", handler.Index)
	r.GET("/contact", handler.Contact)
	r.GET("/curriculum-vitae", handler.Cv)
	r.GET("/cwd", handler.Cwd)
	r.GET("/health", handler.Health)
	r.GET("/home", handler.Index)
	r.GET("/whoami", handler.Whoami)

	port := os.Getenv("PORT")
	if port != "8080" {
		port = "8080"
		log.Printf("Defaulting to port %s.\n", port)
	}

	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatal(err)
	}
}

