package main

import "fmt"

func main() {
	var a int

	fmt.Print("berat parsel (gram): ")
	fmt.Scan(&a)

	b := a / 1000
	c := a % 1000
	d := b * 10000

	fmt.Printf("detail berat: %v kg + %v gr", b, c)
	fmt.Println(" ")

	if b > 10 {
		e := 0
		fmt.Printf("detail harga: Rp.%v + Rp. %v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	} else if b < 10 && c >= 500 {
		e := c * 5
		fmt.Printf("detail harga: Rp.%v + Rp.%v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	} else {
		e := c * 15
		fmt.Printf("detail harga: Rp.%v + Rp.%v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	}
}
