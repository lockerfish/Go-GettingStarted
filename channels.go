package main

import (
	"runtime"
)

func main() {

	// channels deal with shared memory problems on
	// concurrent applications by handle the synchronization
	// and ownership of messages

	// abcGen prints out orderly while
	// taking advantage of multiple processors
	runtime.GOMAXPROCS(8)

	ch := make(chan string)
	doneCh := make(chan bool)

	go abcGen(ch)
	go printer(ch, doneCh)

	println("This comes first!")

	// THIRD TRY REMOVED
	// time.Sleep(100 * time.Millisecond)
	<-doneCh
}

// FIRST TRY
// func abcGen(ch chan string) {
// 	for i := byte('a'); i <= byte('z'); i++ {
// 		ch <- string(i)
// 	}
// }
// func printer(ch chan string) {
// 	for {
// 		println(<-ch)
// 	}
// }

// SECOND TRY
// func abcGen(ch chan string) {
// 	for i := byte('a'); i <= byte('z'); i++ {
// 		ch <- string(i)
// 	}
// 	close(ch)
// }
// func printer(ch chan string) {
// 	for l := range ch {
// 		println(l)
// 	}
// }

// THIRD TRY
func abcGen(ch chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		ch <- string(i)
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}

	doneCh <- true
}
