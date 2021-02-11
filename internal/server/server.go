package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Form struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	Message string `json:"message" xml:"message" form:"message" query:"message"`
}

func Run() {

	r := fiber.New()


	r.Static("/", "../../web/app/dist")

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../../web/app/dist/index.html")
	})

	r.Get("/curriculum-vitae", func(c *fiber.Ctx) error {
		return c.SendString("CV page (not a part of the Vue app, Go page)")
	})

	r.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Site Reliability Dashboard (not a part of the Vue app, Go page)")
	})

	r.Post("/sendMail", func(c *fiber.Ctx) error {
		f := new(Form)

		if err := c.BodyParser(f); err != nil {
			log.Println(err)
			c.Status(500).Send([]byte("failed"))
			return err
		}

		email(f.Name, f.Email, f.Message)

		return c.Send([]byte(f.Email))
	})

	log.Fatal(r.Listen(":8080"))
}