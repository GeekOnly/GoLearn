package main

import (
	"fmt"
	"time"
)

func task(c, quit chan string) {
	for {
		select { // select จัดการขูอมูลผ่าน Channel
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
		default: // Running Forever
			fmt.Println("waiting...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() { // Waiting for result from task
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//quit <- "end" test no end
	}()

	task(c, quit)
}
