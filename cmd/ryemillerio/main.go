package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	r := fiber.New()
	r.Static("../../web", "dist")

	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("../../web/dist/index.html")
	})

	r.Get("/curriculum-vitae", func(ctx *fiber.Ctx) error {
		return ctx.SendString("CV page.")
	})

	r.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Site Reliability Dashboard.")
	})

	log.Fatal(r.Listen(":3000"))
}