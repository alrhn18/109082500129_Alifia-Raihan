package main

import "fmt"

type arr_balita [100]float64

func cari_min (arr_berat arr_balita, n int)float64{
	min := arr_berat[0]

	for i := 1; i < n; i++{
		if arr_berat[i] < min{
			min = arr_berat[i]
		}
	}
	return min
}

func cari_max (arr_berat arr_balita, n int)float64{
	max := arr_berat[0]

	for i := 1; i < n; i++{
		if arr_berat[i] > max {
			max = arr_berat[i]
		}
	}

	return max
}

func rata(arr_berat arr_balita, n int)float64{
	total := 0.0
	for i := 0; i < n; i++{
		total = total + arr_berat[i]
	}

	return total / float64(n)
}

func main () {
	var data arr_balita
	var n int
	var min, max, mean float64

	fmt.Print("banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		fmt.Printf("berat bayi ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	min = cari_min(data, n)
	max = cari_max(data, n)
	mean = rata(data, n)

	fmt.Printf("berat balita minimum: %.2f kg", min)
	fmt.Println("")
	fmt.Printf("berat balita maximum: %.2f kg", max)
	fmt.Println("")
	fmt.Printf("rata rata berat balita: %.2f kg", mean)
}