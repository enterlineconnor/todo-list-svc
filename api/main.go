package main

import (
    "log"
	"api/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	routes.TodoNotes(app)

    log.Fatal(app.Listen(":3000"))
}
