// Generic calls for TODO app

package routers

import (
	"strings"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func TodoNotes(app *fiber.App, db *sql.DB) {
    
	// Return generic hello
	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	// Return custom hello
	app.Get("/user/+", func(c *fiber.Ctx) error {
		name := c.Params("+")
		return_statement := fmt.Sprintf("Hello, %s", strings.Title(name)) 
		return c.SendString(return_statement)
	})
}
