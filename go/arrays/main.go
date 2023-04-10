package main

import "fmt"

func main() {

	// declared array
	// var array_name = [size]type{values}

	// fixed size
	var array1 = [5]int32{}              // not initialize
	array2 := [5]int32{1, 5}             // partial initialize
	array2_5 := [5]int32{1, 5, 3, 4, 5}  // full initialize
	array2_5_1 := [5]int32{1: 10, 2: 40} // specific index initialize
	// dynamic size
	array3 := [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array2_5)
	fmt.Println(array2_5_1)
	fmt.Println(array3)

	// change value && access value

	array1[0] = 10
	array2[1] = 20
	array3[2] = 30

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	// length of array
	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3))

	// slice

	// if this is array (var array_name = [size]type{values})
	// this is slice (var slice_name = []type{values})

	var slice1 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of slice
	fmt.Println(cap(slice1)) // capacity of slice

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// we can also create slice from array
	slice3 := array3[1:3]

	fmt.Println(slice3)

	// proper way to create slice
	// name := make([]type, length, capacity)
	slice4 := make([]int32, 5, 10)

	fmt.Println(slice4)

	slice4 = append(slice4, 1, 2, 3, 4, 5) // append value to slice

	fmt.Println(slice4)

	slice4 = append(slice4, 6, 7, 8) // append more value to slice

	fmt.Println(slice4)

	slice4 = append(slice4, slice3...) // append slice to slice

	fmt.Println(slice4)

	// we can copy slice to another slice

	slice5 := make([]int32, len(slice4))

	copy(slice5, slice4)

	fmt.Println(slice5)

}
