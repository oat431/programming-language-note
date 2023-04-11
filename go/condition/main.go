package main

import "fmt"

func main() {
	a := 4
	b := 7
	c := 1

	// simple bubble sort

	if a > b {
		temp := a
		a = b
		b = temp
	}

	if b > c {
		temp := b
		b = c
		c = temp
	}

	if a > b {
		temp := a
		a = b
		b = temp
	}

	fmt.Println(a, b, c)

	evenNumber := 12

	if evenNumber%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	grade := 56

	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	year := 2023

	// nested condition
	if year%400 != 0 {
		if year%4 == 0 && year%100 != 0 {
			fmt.Println("leap year")
		} else {
			fmt.Println("not leap year")
		}
	} else {
		fmt.Println("leap year")
	}

	// switch case
	day := 9

	switch day {
	case 0, 1:
		fmt.Println("Monday")
	case 2, 9:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

}
