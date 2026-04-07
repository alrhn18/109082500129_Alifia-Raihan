package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var a int

	fmt.Scan(&a)
	
	for i := 0; i <= a; i++ {
		hasil := fibonacci(i)
		fmt.Print(hasil, " ")
	}
}