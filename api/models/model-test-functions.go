// Test functions for canned Backend examples

package animals

import (
	"database/sql"
)
type Animal struct {
	species string `json:"species"`
	name string `json:"name"`
	age int `json:"age"`
	is_dangerous bool `json:"is_dangerous"`
}

type M map[string]interface{}

func GetAllAnimals(db *sql.DB) ([]M, error) {
	query :=`
	SELECT
		species,
		name,
		age,
		is_dangerous
	FROM _animals
	`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
	
	
	var animals []M

	for rows.Next() {
        var animal Animal
        if err := rows.Scan(&animal.species, &animal.name, &animal.age,
            &animal.is_dangerous); err != nil {
				return nil, err
        }

		animal_map := M{
			"species": animal.species,
			"name": animal.name,
			"age": animal.age,
			"is_dangerous": animal.is_dangerous,
		}
		animals = append(animals, animal_map)
	}
	return animals, nil

}