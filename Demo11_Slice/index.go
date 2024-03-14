package main

import "fmt"

func main() {
	var numbers1 = make([]int, 3, 5) // 3 len, 5 cap
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	numbers1 = append(numbers1, 3)
	showSLice(numbers1)

	// ถ้า Append เข้าไปแล้วเกินจำนวนจะเพิ่มขนาด array เป็น 2 เท่า
	var numbers2 []int
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)
	numbers2 = append(numbers2, 4)
	numbers2 = append(numbers2, 5)
	showSLice(numbers2)
}

func showSLice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
