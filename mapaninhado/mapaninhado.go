package main

import "fmt"

func main() {
	// maps podem ter maps com valores
	mapDeMaps := map[string]map[string]float64{
		"A": {
			"Andre": 3000,
		},
		"F": {
			"Fabricio": 4000.00,
		},
	}

	fmt.Println(mapDeMaps)

	delete(mapDeMaps, "F")

	// após remoção de item
	fmt.Println(mapDeMaps)
}
