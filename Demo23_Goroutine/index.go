package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run 1 something :", i)
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run 2 something :", i)
	}
}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
