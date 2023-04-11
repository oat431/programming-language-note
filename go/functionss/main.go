package main

import "fmt"

// type 1 function no return, no parameter
func sayHello() {
	fmt.Println("Hello")
}

// type 2 function with no return, required parameter
func isEven(num int32) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

// type 3 function with return, no parameter
func getHello() string {
	return "Hello"
}

// type 4 function with return, required parameter
func isPrime(num int32) bool {
	if num <= 1 {
		return false
	}
	for i := int32(2); i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// recursive function
func factorial(num int32) int32 {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func fibonacci(num int32) int32 {
	if num == 1 || num == 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	num := int32(10)

	sayHello()
	isEven(num)
	fmt.Println(getHello())

	if isPrime(num) {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}

	fmt.Println("factorial of", num, "is", factorial(num))

	fmt.Println("fibonacci of", num, "is", fibonacci(num))
}
