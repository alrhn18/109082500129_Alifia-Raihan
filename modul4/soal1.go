package main

import "fmt"

func faktorial(n int)int {
	hasil := 1
	for i := 1; i <= n; i++{
		hasil = hasil * i
	}

	return hasil
}

func P(n, r int)int{
	return faktorial(n) / faktorial(n-r)
}

func C(n, r int)int{
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main () {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(P(a, c), C(a, c))
	fmt.Println(P(b, d), C(b, d))
}