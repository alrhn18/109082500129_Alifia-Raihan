package main 

import "fmt"

type kelinci struct{
	berat_kelinci float64
	banyak_kelinci int
}

type arr_kelinci [1000]kelinci

func brt_1(k arr_kelinci, n int)float64{
	var terkecil float64 = k[0].berat_kelinci
	var j int = 1

	for j < n{
		if terkecil > k[j].berat_kelinci{
			terkecil = k[j].berat_kelinci 
		}
		j++
	}
	return terkecil
}

func brt_2(k arr_kelinci, n int)float64{
	var terbesar float64 = k[0].berat_kelinci
	var j int = 1
	
	for j < n{
		if terbesar < k[j].berat_kelinci{
			terbesar = k[j].berat_kelinci
		}
		j++
	}

	return terbesar
}

func main () {
	var data arr_kelinci
	var n int 

	fmt.Print("banyak kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++{
		data[i].banyak_kelinci = i + 1
		fmt.Printf("berat kelinci %d: ", data[i].banyak_kelinci)
		fmt.Scan(&data[i].berat_kelinci)
	}

	min := brt_1(data, n)
	max := brt_2(data, n)

	fmt.Printf("berat terkecil: %.2f", min)
	fmt.Println(" ")
	fmt.Printf("berat terbesar: %.2f", max)
}