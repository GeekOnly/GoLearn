package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}

	if someValue > 10 || someValue < 2 {
		fmt.Println("someValue > 10 || someValue < 20")
	} else {
		fmt.Println("NOT someValue > 10 || someValue < 20")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fnSwitchCase()
}

func doSomething() string {
	// do something
	return "ok"
}

func fnSwitchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("0 is selected")
		break
	case 1:
		fmt.Println("1 is selected")
		break
	case 2:
		fmt.Println("2 is selected")
		break
	default:
		fmt.Println("default is selected")
		break
	}
}
