package main

import "fmt"

func main() {
	// Arrays
	// - Fixed Length
	// - Indexable
	// - Same Type
	// - Contiguous in Memory
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	intArr[1] = 123

	// Access memory locations usig &
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Other ways to initialize arrays
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3)

	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr4)

	// Slices are just arrays
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// But cannot access ranges that don't have values in a slice
	// fmt.Println(intSlice[4])

	// Assign multiple values to a slice using the spread operator
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// using make to define length and capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

}
