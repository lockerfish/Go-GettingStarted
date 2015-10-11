package main

import (
	"runtime"
	"time"
)

func main() {

	// thread-like construct

	// when the application has multiple things todo

	// concurrent
	// single CPU available
	// by default goroutines run single CPU

	// parallel
	// multiple CPU available
	// we use GOMAXPROCS to allow more CPUs
	runtime.GOMAXPROCS(8)

	// abcGen()

	// nothing will print because main function
	// here is done before abcGen completes it just
	// have enough time to schedule abcGen to run.
	// abcGen runs concurrently
	go abcGen()

	println("This comes firsts!")

	// it will print now when we pause the program
	// the abcGen function is run by goroutines
	time.Sleep(100 * time.Millisecond)

}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		// println(string(l))
		go println(string(l))
	}
}
