package main

func show(a, b int) int {
	c := a + b
	return c
}

func test(i, j int) (int, int) {
	a := i + j
	b := i - j
	return a, b
}

func main() {

	println(show(1, 2))
	test(1, 2)
}
