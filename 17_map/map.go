package main

import "fmt"

func main() {
	nameToAgeMap := map[string]int{"subha": 26, "deep": 28, "sunny": 24}
	fmt.Println(nameToAgeMap)

	nameToAgeMap = make(map[string]int)
	fmt.Println(nameToAgeMap)

	nameToAgeMap["subhadeep"] = 27
	nameToAgeMap["subham"] = 25
	fmt.Println(nameToAgeMap)

	fmt.Println("Age of Subhadeep is:", nameToAgeMap["subhadeep"])

	nameToAgeMap["subham"] = 29
	fmt.Println("Age of Subham is:", nameToAgeMap["subham"])

	fmt.Println()
	for key, val := range nameToAgeMap {
		fmt.Println(key, "=", val)
	}

	delete(nameToAgeMap, "subham")
	fmt.Println(nameToAgeMap)

	val, isPresent := nameToAgeMap["subham"]
	fmt.Printf("subham > Value: %d, Present? %t\n", val, isPresent)

	points := map[string]point{"point1": {10, 15}, "point2": {5, 8}, "point3": {1, 25}}
	fmt.Printf("Map of Points: %v\n", points)
}

type point struct {
	x int
	y int
}
