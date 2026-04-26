package main

func main() {
	a := 0
	a--
	println(a)
	a = a >> 1
	println(a)

	b := 1.0
	b += 1.1
	b += 0.9
	b = b + 1.9
	b = b - 100
	println(b)

	s := "Hello "
	s += "World"
	println(s)
}