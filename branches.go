package main

func main() {

	// if-else branches
	foo := 1
	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	// if-else branches
	// initializing inside if statement
	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	// switch
	bar := 1
	switch bar {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("yay!")
	}

	switch {
	case bar == 1:
		println("bar is one")
	case bar > 2:
		println("bar is greater than two")
	}

	// switch
	// initializing inside statement
	switch bar := 2; bar {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("yay!")
	}
}
