package main

import (
	"fmt"
	"practice/moreStructure/models"
)

func main() {
	fmt.Println("This is Struct Practice")
	sample := models.Animal{
		Name: "Dog",
		Age:  10,
	}
	fmt.Println(models.ToString(sample))
}
