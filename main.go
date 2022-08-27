package main

func Add(a, b int) int {
	return a + b
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
}
