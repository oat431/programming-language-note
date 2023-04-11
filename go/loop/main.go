package main

import "fmt"

func main() {
	sum := 0
	// for loop
	for i := 1; i <= 100; i++ {
		if i%16 == 0 {
			continue
		}
		if i == 98 {
			break
		}
		sum += i
	}

	fmt.Println(sum)

	// nested for loop
	sum2 := 0
	for i := 1; i <= 10; i++ {
		if i == 7 {
			continue
		}
		for j := 1; j <= 10; j++ {
			if i == 5 && j == 5 {
				break
			}
			sum2 += j
		}
	}

	fmt.Println(sum2)

	// range keyword

	var fruits = []string{"apple", "banana", "grape", "melon"}
	for no, fruit := range fruits {
		fmt.Printf("%v\t%v\n", no, fruit)
	}

	// while loop style
	fmt.Println("3X+1 Problem")

	num := 27
	temp := num
	step := 0

	for num > 1 {
		step++
		fmt.Println(num)
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
	}

	fmt.Println("for", temp, "take", step, "step")

}
