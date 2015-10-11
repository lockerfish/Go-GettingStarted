package main

func main() {
	// passing 'foo' fails
	// we need a constructor on the enhancedMessagePrinter
	// mp := enhancedMessagePrinter{"foo"}
	mp := enhancedMessagePrinter{}
	mp.message = "foo"
	mp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
