package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	r := fiber.New()
	r.Static("/", "web/app/dist")

	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("web/app/dist/index.html")
	})

	r.Get("/curriculum-vitae", func(ctx *fiber.Ctx) error {
		return ctx.SendString("CV page (not a part of the Vue app, Go page)")
	})

	r.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Site Reliability Dashboard (not a part of the Vue app, Go page)")
	})

	log.Fatal(r.Listen(":8080"))
}