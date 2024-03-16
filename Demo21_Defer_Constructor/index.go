package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer printFinished(i)
		// defer จาก 9 - 0 Stack Buffer
	}
	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}
}

func printFinished(i int) {
	fmt.Println("Finish: ", i)
}
