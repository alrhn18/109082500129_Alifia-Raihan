package main 

import "fmt"

func faktor (a, pembagi int){
	if pembagi > a{
		return
	}

	if a % pembagi == 0{
		fmt.Print(pembagi, " ")
	}

	faktor(a, pembagi + 1)
}

func main () {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}