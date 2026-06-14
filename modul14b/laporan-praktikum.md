# <h1 align="center">Laporan Praktikum Modul 14 part 2 </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Mengurutkan Angka dari Kecil ke Besar dan Mencari Jarak Antara Angka

### 1. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

#### soal1.go

```go
package main

import "fmt"

const nMax = 1000

type arrInt [nMax]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func checkJarak(T *arrInt, n int) string {
	var jarak, k int
	var tetap bool
	jarak = T[1] - T[0]
	tetap = true
	k = 2
	for k < n && tetap {
		if T[k]-T[k-1] != jarak {
			tetap = false
		}
		k = k + 1
	}
	if tetap {
		return fmt.Sprintf("Data berjarak %d", jarak)
	}
	return "Data berjarak tidak tetap"
}

func main() {
	var arr arrInt
	var n, x, k int

	fmt.Scan(&x)
	for x >= 0 {
		arr[n] = x
		n = n + 1
		fmt.Scan(&x)
	}

	insertionSort(&arr, n)

	for k < n {
		if k > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[k])
		k = k + 1
	}
	fmt.Println()
	fmt.Println(checkJarak(&arr, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal1.jpg)

##### Penjelasan
Kode ini untuk mengurutkan angka dari kecil ke besar dengan metode insertion sort dan memeriksa jarak antar angka, menggunakan array bermuatan 1000.

Pada func insertionSort untuk mengurutkan angka dari kecil ke besar (ascending) menggunakan insertion sort. Dengan cara mengambil elemen ke-i dan diumpamakan sebagai temp, lalu membandingkan dengan elemen sebelumnya dan menggeser elemen yang lebih besar ke kanan hingga ditemukan posisi yang tepat untuk temp.

Pada func checkJarak untuk menghitung dan mengecek jarak antar angka tersebut. Dengan cara menghitung selisih elemen indeks 1 dan indeks 0 sebagai acuan, lalu membandingkan selisih tiap elemen berikutnya. Jika ada satu selisih yang berbeda, maka akan memberi output "data berjarak tidak tetap", tetapi jika sama semua akan memberi output "data berjarak x" dimana x adalah selisih antar indeks.

Pada func main program membaca input angka satu per satu dan berhenti ketika menerima angka negatif. Kemudian memanggil func insertionSort untuk mengurutkan dan mencetak hasilnya, lalu memanggil func checkJarak untuk mengecek jarak antar elemen.

## Mencari Data Buku Terfavorit, Urutan Judul Buku dengan Rating Tertinggi, dan Data Buku sesuai Rating yang Dicari

### 2.  Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Dengan masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. Dan keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### soal2.go

```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func daftarkanBuku() {
	fmt.Scan(&nPustaka)
	for i := 0; i < nPustaka; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func cetakTerfavorit(p DaftarBuku, n int) {
	if n == 0 {
		return
	}
	idx := 0
	for i := 1; i < n; i++ {
		if p[i].rating > p[idx].rating {
			idx = i
		}
	}
	fmt.Println(p[idx].judul, p[idx].penulis, p[idx].penerbit, p[idx].tahun)
}

func urutBuku(p *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && p[j].rating < key.rating {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}
}

func cetak5Terbaru(p DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(p[i].judul)
	}
}

func cariBuku(p DaftarBuku, n int, r int) {
	kn, kr := 0, n-1
	hasil := -1
	for kn <= kr && hasil == -1 {
		tgh := (kn + kr) / 2
		if p[tgh].rating == r {
			hasil = tgh
		} else if p[tgh].rating < r {
			kr = tgh - 1
		} else {
			kn = tgh + 1
		}
	}
	if hasil == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		fmt.Println(p[hasil].judul, p[hasil].penulis, p[hasil].penerbit, p[hasil].tahun, p[hasil].eksemplar, p[hasil].rating)
	}
}

func main() {
	var r int
	daftarkanBuku()

	cetakTerfavorit(pustaka, nPustaka)

	urutBuku(&pustaka, nPustaka)

	cetak5Terbaru(pustaka, nPustaka)

	fmt.Scan(&r)
	cariBuku(pustaka, nPustaka, r)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal2.jpg)

##### Penjelasan
Kode ini untuk mencari data buku terfavorit, urutan judul buku dengan rating tertinggi, dan data buku sesuai rating yang dicari menggunakan insertion sort, sequential search, dan binary search.

Pada struct Buku untuk menyimpan data buku yaitu berupa id, judul, penulis, penerbit, eksemplar, tahun, dan rating. Dengan array bermuatan 7919.

Pada func daftarkanBuku untuk menginput id, penulis, penerbit, eksemplar, tahun dan rating.

Pada func cetakTerfavorit untuk mencari buku dengan rating tertinggi menggunakan sequential search sebelum array diurutkan.

Pada func urutBuku untuk mengurutkan buku dari rating tertinggi ke rating terkecil (descending) menggunakan insertion sort.

Pada func cetak5Terbaru untuk mencetak urutan 5 buku dengan rating tertinggi. Variabel limit untuk mengantisipasi jika jumlah buku yang diinput kurang dari 5, maka nanti akan mencetak urutan buku sesuai yang diinput.

Pada func cariBuku untuk mencari buku berdasarkan rating menggunakan binary search. Dengan kn dan kr jadi batas pencarian. Setiap iterasi akan mencari indeks tengah, lalu membandingkan rating yang ingin dicari.

Pada func main untuk menjalankan dan memanggil seluruh fungsi secara berurutan, yaitu daftarkanBuku untuk menginput data buku, cetakTerfavorit untuk mencetak buku dengan rating tertinggi, urutBuku untuk mengurutkan buku, cetak5Terbaru untuk mencetak urutan buku dari ratin tertinggi ke rating terkecil, dan cariBuku untuk mencari buku dengan rating yang diinput.