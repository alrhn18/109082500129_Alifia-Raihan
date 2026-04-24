# <h1 align="center">Laporan Praktikum Modul 10 </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Mencari Minimal dan Maximal Berat Kelinci 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya nN bilangan rill berikutnya adlaah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan yang menyatakan berat kelinci terkecil dan terbesar.

#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul10/output/output-soal1.jpg)

##### Penjelasan
code diatas untuk mencari berat kelinci terkecil dan terbesar menggunakan array. 

func brt_1 untuk mencari berat kelinci terkecil, func brt_2 untuk mencari berat kelinci terbesar. pada func main akan menginput banyak kelinci, dan berat kelinci. lalu di func main juga akan memberi output berat kelinci terkecil dan terbesar.


## (belum di kerjakan -soal lele-)

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul10/output/output-soal2.jpg)

##### Penjelasan
code diatas untuk . 

-belum-


## Mencari Min, Max, dan Rata Rata Berat Badan Balita

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul10/output/output-soal3.jpg)

##### Penjelasan
code diatas untuk mencari min, max, dan rata rata berat badan balita


func cari_min untuk mencari berat badan balita terkecil. func cari_max untuk mencari berat badan balita terbesar. func rata untuk mencari rata rata berat badan balita. di func main untuk menginput benyak balita yang akan ditimbang, dan menginput berat badan balita. di func main juga akan memberikan output berat badan terkecil, berat badan terbesar, dan rata rata.