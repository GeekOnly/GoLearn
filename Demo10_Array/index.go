package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3}
	var array2 = []int{1, 2, 3, 4, 5, 6}
	var array3 [3]string

	for _, item := range array1 {
		fmt.Println(item, ",")
	}

	for _, item := range array2 {
		fmt.Println(item, ",")
	}

	array3[0], array3[1], array3[2] = "Android", "IOS", "Mac"

	for _, item := range array3 {
		fmt.Println(item, ",")
	}
}
