package main 

import "fmt"

type lele struct{
	berat float64
	nomor int
}

type arrlele [1000]lele

func hitung_wadah(data arrlele, x, y int){
	jumlah_wadah := x / y
	
	if x % y != 0{
		jumlah_wadah++
	}

	// total
	var total[1000]float64
	var wadah int

	j := 0
	for j < x{
		wadah = j / y
		total[wadah] += data[j].berat
		j++
	}

	i := 0
	for i < jumlah_wadah{
		if i > 0{
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", total[i])
		i++
	}
	fmt.Println()

	// rata-rata
	total_semua := 0.0
	k := 0

	for k < jumlah_wadah{
		total_semua += total[k]
		k++
	}

	rata_rata := total_semua / float64(jumlah_wadah)
	fmt.Printf("%.2f", rata_rata)
}

func main () {
	var data arrlele
	var x, y int

	fmt.Print("banyak ikan dan isi wadah: ")
	fmt.Scan(&x, &y)

	i := 0
	for i <x{
		data[i].nomor = i +1
		fmt.Printf("berat ikan ke-%d: ", data[i].nomor)
		fmt.Scan(&data[i].berat)
		i++
	}

	hitung_wadah(data, x, y)
}

