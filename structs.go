package main

import (
	"fmt"
)

func main() {

	// creating objects in local execution stack

	// create myStruct object
	bar := myStruct{}
	bar.myField = "bar"
	fmt.Println(bar.myField)

	// passing field value, must be in same order as
	// defined on myStruct
	foo := myStruct{"foo"}
	fmt.Println(foo.myField)
	fmt.Println(foo)

	// often better to create objects
	// on the heap instead, like this
	heap := new(myStruct)
	heap.myField = "heap"
	fmt.Println(heap.myField)
	// heap is in fact a memory address
	// printing value on console
	// not need to de-reference pointer
	// much easier to work with referenced data type
	fmt.Println(heap)

	// same but without using "new" keyword
	// now we can pass value like before
	xoom := &myStruct{"xoom"}
	fmt.Println(xoom.myField)
	fmt.Println(xoom)
}

type myStruct struct {
	myField string
}
