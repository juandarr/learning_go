package main

import "fmt"

func fib(n int) int {
	var a int = 0
	var b int = 1
	var tmp int
	for i := 0; i < n; i++ {
		tmp = b
		b = a + b
		a = tmp
	}
	return b
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("Term %v of fib is %v\n", i, fib(i))
	}
}
