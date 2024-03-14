package main

import "fmt"

func main() {
	fn1()
	fn2("XS Game Studio in Germany")
	fn3("Farcry", 55)
}

func fn1() {
	fmt.Println("XSGame Display")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}
