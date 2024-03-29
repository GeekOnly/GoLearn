package main

import "fmt"

func task(c, quit chan string) {
	for {
		select { // select จัดการขูอมูลผ่าน Channel
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
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
		quit <- "end"
	}()

	task(c, quit)
}
