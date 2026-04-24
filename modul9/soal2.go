package main 

import "fmt"
import "math"


func main () {
	var n, x, hapus_index, cari int
	var varsum float64

	fmt.Print("n: ")
	fmt.Scan(&n)

	// membuat array
	arr := make ([]int, n)
	for i:= 0; i <= n-1; i++{
		fmt.Printf("arr[%d]: ", i)
		fmt.Scan(&arr[i])
	}

	// semua elemen
	fmt.Println("semua:", arr)

	// elemen yang keluar jika index ganjil 
	fmt.Print("elemen dari indeks ganjil: ")
	for i := 0; i < len(arr); i++{
		if i % 2 != 0{
			fmt.Print(arr[i], " ")
		}
	}

	// elemen yang keluar jika index genap
	fmt.Println(" ")
	fmt.Print("elemen dari indeks genap: ")
	for i:= 0; i < len(arr); i++{
		if i % 2 == 0{
			fmt.Print(arr[i], " ")
		}
	}

	// untuk mencari elemen dari indeks kelipatan x
	fmt.Println(" ")
	fmt.Print("x: ")
	fmt.Scan(&x)
	fmt.Printf("elemen pada indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++{
		if i % x == 0{
			fmt.Print(arr[i], " ")
		}
	}

	// untuk menghapus elemen pada indeks tertentu
	fmt.Println(" ")
	fmt.Print("hapus indeks: ")
	fmt.Scan(&hapus_index)
	arr = append(arr[:hapus_index], arr[hapus_index+1:]...)
	fmt.Println("elemen setelah dihapus:", arr)

	// untuk rata rata
	sum := 0
	for i := 0; i < len(arr); i++{
		sum = sum + arr[i]
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Printf("rata rata: %.2f", mean)

	// untuk simpangan baku
	fmt.Println(" ")
	for i := 0; i < len(arr); i++ { 
		varsum = varsum + math.Pow(float64(arr[i])-mean, 2)
	}
	sb := math.Sqrt(varsum/float64(len(arr)))
	fmt.Printf("simpangan baku: %.2f", sb)

	// untuk cari frekuensi
	fmt.Println(" ")
	fmt.Print("cari bilangan: ")
	fmt.Scan(&cari)
	hitung := 0
	for i := 0; i < len(arr); i++{
		if arr[i] == cari{
			hitung = hitung + 1
		}
	}
	fmt.Printf("frekuensi %d: %d kali", cari, hitung)
}