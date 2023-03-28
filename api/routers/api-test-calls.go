// Test calls for canned CRUD examples

package routers

import (
	"encoding/json"
	"fmt"
	"database/sql"
	
	"api/models"

	"github.com/gofiber/fiber/v2"
)
  
func ApiTestCalls(app *fiber.App, db *sql.DB) {

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
		person := string(jsonData)
		return c.SendString(person)
    })

	app.Get("/animals", func(c *fiber.Ctx) error {
		animals_list, err := animals.GetAllAnimals(db)

		if err != nil {
			fmt.Printf("Something went wrong with animal retrieval...")
			return err
		}

		return c.Status(fiber.StatusOK).JSON(animals_list)

	  })

}

