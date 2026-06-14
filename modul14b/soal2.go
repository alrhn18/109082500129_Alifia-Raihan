package main 

import "fmt"

type arrKerabat [1000000]int

//untuk menghitung ganjil
func rumahkerabat_ganjil(T *arrKerabat, n int){
	var a, i, j, idx_min int

	i = 1
	for i <= n -1{
		idx_min = i -1
		j = i 
		for j < n{
			if T[idx_min] > T[j]{
				idx_min = j
			}
			j = j + 1
		}
		a = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = a
		i = i + 1
	}
}

//untuk menghitung genap
func rumahkerabat_genap(T *arrKerabat, n int){
	var a, i, j, idx_min int

	i = 1
	for i <= n -1{
		idx_min = i -1
		j = i 
		for j < n{
			if T[idx_min] < T[j]{
				idx_min = j
			}
			j = j + 1
		}
		a = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = a
		i = i + 1
	}
}

func main () {
	var n, m, x int
	var genap, ganjil, gabung arrKerabat

	fmt.Scan(&n)

	d := 1
	for d <= n{
		fmt.Scan(&m)
		gj := 0 //ganjil
		gn := 0 //genap

		k := 0
		for k < m{
			fmt.Scan(&x)
			if x % 2 != 0{
				ganjil[gj] = x
				gj = gj + 1
			}else if x % 2 == 0{
				genap[gn] = x
				gn = gn + 1
			}
			k++
		}
		
		//urutin dulu dari ganjil (kecil ke besar) lanjut ke genap (kecil ke besar)
		if gj > 0{
			rumahkerabat_ganjil(&ganjil, gj)
		}

		if gn > 0{
			rumahkerabat_genap(&genap, gn)
		}

		//untuk gabungin ganjil dan genap, untuk di cetak
		k = 0
		for k < gj{
			gabung[k] = ganjil[k]
			k++
		}

		k = 0
		for k < gn{
			gabung[gj+k] = genap[k]
			k++
		}
		total := gj + gn

		//untuk mencetak
		for k = 0; k < total; k++ {
			if k < total-1 {
				fmt.Printf("%d ", gabung[k])
			} else {
				fmt.Printf("%d\n", gabung[k])
			}
		}
		d++
	}	
}