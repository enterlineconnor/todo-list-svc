// Test calls for canned CRUD examples

package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"fmt"
)

func ApiTestCalls(app *fiber.App) {
    

	app.Get("/api-test", func(c *fiber.Ctx) error {
        data := map[string]interface{}{
			"firstName": "John",
			"lastName": "Doe",
			"age": 23,
			"title": "Full Stack Developer",
		}
	
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return err
		}
		fmt.Printf("%s\n", jsonData)
		blah := string(jsonData)
		return c.SendString(blah)
    })

}
