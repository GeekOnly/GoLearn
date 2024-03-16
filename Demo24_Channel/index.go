package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1 // Send a signal
	fmt.Println("step 1")
	fmt.Println(<-ch)

	ch <- 2 // Send a signal
	fmt.Println("step 2")
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)
}
