package main

import "fmt"

func main() {
	var anyType = "This is any type variable"

	fmt.Println(anyType)

	numberType := 622115039

	fmt.Println("this is number type : ", numberType)

	var a string
	var b int
	var c float32
	var d bool

	fmt.Println("default value of : ", a, b, c, d)

	const PI = 3.14

	fmt.Println("this is constant value : ", PI)

	const G float32 = 9.8

	fmt.Println("this is constant value : ", G)

	// decalre a variable use :=
	// change a variable value use =

	// foo := 32 is equivalent to var foo int = 32

	// important note for integer type

	var normalInt int = 431 // default 32 bit or 64 bit depend on system so you need to declare clearly

	var int8Type int8 = 127                   // 8 bit
	var int16Type int16 = 32767               // 16 bit
	var int32Type int32 = 2147483647          // 32 bit
	var int64Type int64 = 9223372036854775807 // 64 bit

	// you can also do unsigned integer type
	var normalUint uint = 431                    // default 32 bit or 64 bit depend on system so you need to declare clearly
	var uint8Type uint8 = 255                    // 8 bit
	var uint16Type uint16 = 65535                // 16 bit
	var uint32Type uint32 = 4294967295           // 32 bit
	var uint64Type uint64 = 18446744073709551615 // 64 bit

	fmt.Println(normalInt, int8Type, int16Type, int32Type, int64Type)
	fmt.Println(normalUint, uint8Type, uint16Type, uint32Type, uint64Type)

	// float has 2 type
	// float32 and float64
	var float32Type float32 = 3.14
	var float64Type float64 = 6.28

	fmt.Println(float32Type, float64Type)

}
