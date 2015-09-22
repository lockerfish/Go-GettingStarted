package main

const (
	first  = "the first"
	second = "the second"
	third  = "the third"
)

// iota gives us an auto-incrementing number
const (
	uno = iota
	dos
	tres
)

// constant expression
// bit shiftting (2 raise to 10 times iota power)
const (
	green = 1 << (10 * iota)
	blue
	red
)

func main() {
	println(first)
	println(second)
	println(third)

	println(uno)
	println(dos)
	println(tres)

	println(green)
	println(blue)
	println(red)
}
