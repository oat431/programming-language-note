package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func pointInfo(p Point) {
	fmt.Println("x:", p.x, "y:", p.y)
}

func findDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) - math.Pow((p1.y-p2.y), 2))
}

func main() {
	var p1 Point
	p1.x = 10.0
	p1.y = 20.0
	pointInfo(p1)

	var p2 Point = Point{20.0, 10.0}
	pointInfo(p2)

	fmt.Println("ended of the program")

	fmt.Println("distance:", findDistance(p1, p2))
}
