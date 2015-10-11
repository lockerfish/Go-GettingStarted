package main

func main() {
	message := "Hello, World!"
	sayHelloNoParams()
	sayHelloWithParams(message)

	sayHelloPassingByValue(message)
	println("message still the same: " + message)

	// passing address of message
	sayHelloPassingByRef(&message)
	println("message is NOT the same: " + message)

	sayHelloVariatic("Hello", "Go", "from", "Hendrix")

	// returning a single value
	result := addSingleReturn(1, 3, 5, 9)
	println(result)

	// returning multiple values
	numTerms0, sum0 := addMultiReturns(1, 3, 5, 9)
	println("Added:", numTerms0, "terms for a total of", sum0)

	// using named parameters
	numTerms1, sum1 := addNamedParams(1, 3, 5, 9)
	println("Added:", numTerms1, "terms for a total of", sum1)

	// anonymous function
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}
	numTerms2, sum2 := addFunc(1, 3)
	println("Added:", numTerms2, "terms for a total of", sum2)
}

func sayHelloNoParams() {
	println("Hello")
}

func sayHelloWithParams(message string) {
	println(message)
}

func sayHelloPassingByValue(message string) {
	println(message)

	message = "Hello Go"
}

func sayHelloPassingByRef(message *string) {
	// de-referencing pointer
	println(*message)

	*message = "Hello Go"
}

func sayHelloVariatic(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

// function returning a single value
func addSingleReturn(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}

// function returning multiple values
func addMultiReturns(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

// function with named parameters
func addNamedParams(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
