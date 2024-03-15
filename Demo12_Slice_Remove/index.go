package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	showSLice(numbers)

	// leading removes
	numbers = numbers[1:len(numbers)]
	showSLice(numbers)
	numbers = numbers[1:len(numbers)]
	showSLice(numbers)

	// trailing remove
	numbers = numbers[0 : len(numbers)-1]
	showSLice(numbers)
	numbers = numbers[0 : len(numbers)-1]
	showSLice(numbers)

	// remvoe specifi index
	numbers = removeIndex(numbers, 2)
	showSLice(numbers)
	numbers = removeIndex(numbers, 1)
	showSLice(numbers)
}

func showSLice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
