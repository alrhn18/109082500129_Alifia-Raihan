package main 

import "fmt"

func main () {
	var (
		a, b, c, d string
		i int
		e bool
	)

	e = true
	for i = 1; i <= 5; i++ {
		fmt.Printf("percobaan ke-%d: ", i)
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu"{
			e = false
		}
	}
	fmt.Println("hasil:", e)
}