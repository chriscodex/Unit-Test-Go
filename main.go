package main

func Add(a, b int) int {
	return a + b
}

var m = make(map[int]int, 0)

func fibonacciQuick(n int) int {
	if n < 2 {
		return n
	}

	var f int
	if v, ok := m[n]; ok {
		f = v
	} else {
		f = fibonacciQuick(n-2) + fibonacciQuick(n-1)
		m[n] = f
	}

	return f
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
}
