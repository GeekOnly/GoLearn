package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(msgPointer)

	changeMessage(&msg, "New Message 1")
	fmt.Println(msg)

	changeMessage(msgPointer, "New Message 2")
	fmt.Println(msg)

	changeMessage(msgPointer, "New Message 3")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
