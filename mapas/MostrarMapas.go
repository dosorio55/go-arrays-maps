package mapas

import (
	"fmt"
)

func MostrarMapas() {
	// make is the reserved word to create a map
	// a map is a key value pair, like an object in javascript
	// the first argument is the type of the key and the second the type of the value
	contries := make(map[string]string)

	// we can also add another parameter to the make function to specify the capacity
	// contries := make(map[string]string, 5)

	contries["Mexico"] = "CDMX"
	contries["Argentina"] = "Buenos Aires"

	// we can also create a map with the following syntax
	foods := map[string]string{
		"Mexico":    "Tacos",
		"Argentina": "Asado",
		"Colombia":  "Bandeja Paisa",
		"Peru":      "Ceviche",
		"Chile":     "Empanadas",
	}

	for contry, food := range foods {
		fmt.Printf("The food of %s is %s\n", contry, food)
	}

	// delete a key value pair
	// functions first argument is the map and the second the key
	delete(foods, "Chile")

	// return a certain value of the map
	// the first argument is the map and the second a boolean if exists or not
	food, exists := foods["Colombia"]

	if exists {
		fmt.Printf("The food of Colombia is %s\n", food)
	} else {
		fmt.Println("The food of Colombia does not exists")
	}

	// fmt.Println(contries)
	// // fmt.Println(food)
	MostrarSlices()
}
