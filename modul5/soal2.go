package main 

import "fmt"

func bintang(n int){
	if n == 0{
		return
	}

	fmt.Print("*")
	bintang(n-1)
}

func main () {
	var a int

	fmt.Scan(&a)
	for i := 1; i <= a; i++{
		bintang(i)
		fmt.Println()
	}
}