package main

import "fmt"

var cout int = 0

func main() {
	fmt.Println("Begin")

	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true
	const tmp4 int = 0

	// implicit Declration
	// var tmp5 int = 0
	tmp5 := 0
	tmp6 := "XSGameStudio"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)
	fmt.Println(tmp5)
	fmt.Println(tmp6)

	cout++
	fmt.Println(cout)
	run()
	run()
	run()

}

func run() {
	cout++
	fmt.Println(cout)
}
