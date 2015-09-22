package main

import "fmt"

func main() {
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99

	fmt.Println(myArray)

	// using array initializers
	theArray := [...]int{42, 27, 99}
	fmt.Println(theArray)
	fmt.Println(len(theArray))

	// slice is used instead of array
	mySlice := theArray[:]
	fmt.Println(mySlice)
	mySlice = append(mySlice, 100)
	fmt.Println(mySlice)

	// creating slice
	theSlice := []float32{12., 15., 18.}
	fmt.Println(theSlice)
	fmt.Println(len(theSlice))

	// using make to make slice more efficient
	mSlice := make([]float32, 100)
	mSlice[0] = 12.
	mSlice[1] = 15.
	mSlice[2] = 18.
	fmt.Println(mSlice)
	fmt.Println(len(mSlice))

	// key-value pair
	// map[key]value
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[42] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)
	fmt.Println(myMap[99])
}
