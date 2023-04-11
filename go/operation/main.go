package main

import "fmt"

func main() {
	// arithmetic operation
	a := 10
	a++
	a--
	a += 10
	a -= 8
	a %= 3

	fmt.Println(a)

	fmt.Println(a%2 == 0)

	fmt.Println(a != 3)
	fmt.Println(a == 3)
	fmt.Println(a >= 3)
	fmt.Println(a <= 3)

	fmt.Println(a != 3 && a == 3)
	fmt.Println(a >= 3 || a <= 3)
	fmt.Println(!(a >= 4))

	fmt.Println(a | 3)
	fmt.Println(a & 3)
	fmt.Println(a ^ 3)
	fmt.Println(a << 3)
	fmt.Println(a >> 3)

}
