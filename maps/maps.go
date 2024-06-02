package main

import "fmt"

type floatMapx map[string]float64

func (list floatMapx) output() {
	fmt.Println(list)
}

func main() {
	// Introduction to Maps
	websites := map[string]string{
		"Google": "www.google.com",
		"AWS":    "www.amazon.com",
	}

	fmt.Println(websites)

	// Add to the map
	addToMaps(websites)
	fmt.Println(websites)

	// Delete from the map
	deleteFromMaps(websites)
	fmt.Println(websites)

	// Using make to create a map
	courses := usingMake()
	fmt.Println(courses)

	// Using make with alias type with receiver functions
	courseRating := usingMakeWithAlias()
	courseRating.output()

	// Looping maps
	loopMap(courses)
	fmt.Println()
	loopMap(courseRating)
}

func addToMaps(websites map[string]string) {
	websites["LinkedIn"] = "www.linkedin.com"
}

func deleteFromMaps(websites map[string]string) {
	delete(websites, "Google")
}

func usingMake() map[string]float64 {
	courses := make(map[string]float64, 1)
	courses["Go"] = 19.99
	courses["React"] = 23.99
	courses["NodeJs"] = 49.99
	return courses
}

func usingMakeWithAlias() floatMapx {
	courses := make(floatMapx, 1)
	courses["Go"] = 79.99
	courses["React"] = 93.99
	courses["NodeJs"] = 249.99
	return courses
}

func loopMap(list map[string]float64) {
	// ...
	for key, value := range list {
		fmt.Printf(`Key: %v, Value: %v`, key, value)
		fmt.Println()
	}
}
