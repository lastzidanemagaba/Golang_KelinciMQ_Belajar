package main

import (
	"os"
	"zidane/message"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		panic("Please provide an argument")
	}

	switch arg[1] {
	case "send":
		message.Send()
	case "receive":
		message.Receive()
	default:
		panic("Please provide a valid argument")
	}
}
