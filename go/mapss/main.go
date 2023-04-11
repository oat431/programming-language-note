package main

import "fmt"

func showMap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func main() {
	simpleMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// empty map
	simpleMap2 := make(map[string]int)

	fmt.Println(simpleMap["one"])
	fmt.Println(simpleMap2)

	// delete element from map
	delete(simpleMap, "one")

	// checking if key exists
	val, ok := simpleMap["one"]
	val2, ok2 := simpleMap["two"]

	fmt.Println(val, ok)
	fmt.Println(val2, ok2)

}
