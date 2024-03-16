package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("y = %d", y)
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
	fmt.Println(newInts()) // 2
	fmt.Println(newInts()) // 3
	fmt.Println(newInts()) // 4

	seqString := stringSeq()
	fmt.Println(seqString()) // y = 1
	fmt.Println(seqString()) // y = 2

	fmt.Println(stringSeq()())
}
  