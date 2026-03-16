package main

import "fmt"
import "math"

func jarak (a, b, c, d float64)float64{
	return math.Sqrt(((a-c)*(a-c)) + ((b-d)*(b-d)))
}

func dalam (cx, cy, x, y, r float64)bool{
	return jarak (cx, cy, x, y) <= r
}

func main () {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
	)

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dlm1 := dalam(cx1, cy1, x, y, r1)
	dlm2 := dalam(cx2, cy2, x, y, r2)

	if dlm1 && dlm2 {
		fmt.Println("titik dalam lingkaran 1 dan 2")
	}else if dlm1 {
		fmt.Println("titik dalam lingkaran 1")
	}else if dlm2 {
		fmt.Println("titik dalam lingkaran 2")
	}else{
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}