package main

import "fmt"

func main() {
	var p1 product
	p1.name = "Ardunio"
	p1.price = 100
	p1.stock = 20
	show(p1)

	p2 := createStructData("RasberryPI", 10, 200)
	show(p2)

	update(&p1)
	show(p1)

	//p1 = p1.clear()
	p1 = p1.setDescount(1)
	show(p1)
}

type product struct {
	name  string
	price int
	stock int
}

func update(p *product) {
	p.price = p.price + 100
	p.stock = 10
}

// Function Call
func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDescount(d int) product {
	p.price = p.price - d
	return p
}

func show(p product) {
	fmt.Println(p)
}

func createStructData(name string, price int, stock int) product {
	return product{name, price, stock}
}
