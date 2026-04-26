package main

import "fmt"

func main() {
	var (
		a int
		p *int
	)

	a = 100
	p = &a
	println(*p)
	a = 2
	println(*p)

	b := *p
	a = 1
	fmt.Println(b)
}