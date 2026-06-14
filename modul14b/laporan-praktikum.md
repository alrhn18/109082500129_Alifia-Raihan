# <h1 align="center">Laporan Praktikum Modul 14 (Selection Sort) </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Mengurutkan Nomor Rumah dari Nomor Terkecil ke Terbesar 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 

#### soal1.go

```go
package main

import "fmt"

type arrRumah [1000000]int

func rumahkerabat(T *arrRumah, n int){
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

func main () {
	var n, m int
	var T arrRumah

	fmt.Scan(&n)

	d := 1
	for d <= n{
		fmt.Scan(&m)
		k := 0
		for k < m{
			fmt.Scan(&T[k])
			k = k + 1
		}
		rumahkerabat(&T, m)

		for k := 0; k < m; k++ {
			if k < m-1 {
				fmt.Printf("%d ", T[k])
				} else {
					fmt.Printf("%d\n", T[k])
				}
			}
			d++
		}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal1.jpg)

##### Penjelasan
Kode tersebut untuk mengurutkan nomor rumah dari kecil ke terbesar menggunakan cara sorting. Dengan n adalah banyak daerah kerabat Hercules tinggal dan m adalah banyaknya rumah kerabat di daerah tersebut.

Func rumahkerabat untuk mengurutkan dari nomor terkecil ke nomor terbesar. Bekerja dengan cara mencari nilai terkecil(idx_min) pada data, lalu menukarnya dengan elemen paling kiri menggunakan variabel sementara a.

Pada func main nantinya program akan membaca n daerah. Untuk setiap daerah, dibaca m nomor rumah ke dalam array, kemudian akan memanggil func rumahkerabat dan hasilnya akan dicetak perbaris.


## Mengurutkan Nomor Rumah dari Ganjil (Kecil ke Besar) dilanjut ke Genap (kecil ke besar)

### 2.  Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. 

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal2.jpg)

##### Penjelasan
Kode tersebut untuk mengurutkan nomor rumah dari nomor ganjil (dari nomor terkecil ke nomor besar) ke nomor genap (dari nomor terkecil ke nomor terbesar). Dimana nantinya akan dibedakan func nya untuk menghitung ganjil dan genap.

Func rumahkerabat_ganjil untuk mengurutkan nomor ganjil dari terkecil ke terbesar dengan menggunakan cara sorting. Bekerja dengan cara mencari nilai terkecil (idx_min) lalu menukarnya ke posisi kiri.

Func rumahkerabat_genap untuk mengurutkan nomor genap dari terkecil ke terbesar dengan menggunakan cara sorting. Bekerja dengan mencari idx_min sebagai nilai terbesar, lalu menukarnya ke posisi kanan.

Func main setelah menginnput angka, nantinya angka tersebut langsung dipisahkan ke array ganjil atau genap. Lalu kedua array akan digabung ke array gabung dengan hasil sorting ganjil didepan dan hasil sorting genap dibelakang, lalu dicetak dalam satu baris.


## Mencari Median yang Diminta, Satu Data per Baris

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

#### soal3.go

```go
package main 

import "fmt"

type arrlomba[10000] int

func median(T *arrlomba, n int){
	var i, j, a int

	i = 1
	for i <= n-1{
		j = i
		a = T[j]
		for j > 0 && a < T[j-1]{
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = a
		i = i + 1
	}
}

func main () {
	var skor arrlomba
	var x, n int

	n = 0
	fmt.Scan(&x)
	for x != -5313{
		if x == 0{
			median(&skor, n)
			if n % 2 == 1{
				fmt.Println(skor[n/2])
			}else{
				fmt.Println((skor[n/2-1]+skor[n/2])/2)
			}
		}else{
			skor[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal3.jpg)

##### Penjelasan
Kode tersebut untuk mencari nilai median dari kumpulan data yang sudah terbaca.

Func median untuk mengurutkan nilai yang diinput dari yang terkecil ke terbesar dengan cara sorting. Bekerja dengan cara mengambil satu elemen a, lalu menggeser elemen elemen sebelumya yang lebih besar ke kanan hingga menjadi terurut dari terkecil ke terbesar.

Func main untuk membaca angka yang diinput. Bilangan selain 0 dan -5313 akan disimpan ke array skor. Saat program membaca 0, func median akan dipanggil lalu akan menghitung median. Jika jumlah data ganjil, maka akan diampil elemen yang ditengah, jika jumlah data genap, maka diambil rata rata dua elemen tengah yang dibulatkan kebawah.

Progran ini akan berhenti ketika program membaca -5313


# <h1 align="center">Laporan Praktikum Modul 14 (Insertion Sort) </h1>

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal1b.jpg)

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14b/output/output-soal2b.jpg)

##### Penjelasan
Kode ini untuk mencari data buku terfavorit, urutan judul buku dengan rating tertinggi, dan data buku sesuai rating yang dicari menggunakan insertion sort, sequential search, dan binary search.

Pada struct Buku untuk menyimpan data buku yaitu berupa id, judul, penulis, penerbit, eksemplar, tahun, dan rating. Dengan array bermuatan 7919.

Pada func daftarkanBuku untuk menginput id, penulis, penerbit, eksemplar, tahun dan rating.

Pada func cetakTerfavorit untuk mencari buku dengan rating tertinggi menggunakan sequential search sebelum array diurutkan.

Pada func urutBuku untuk mengurutkan buku dari rating tertinggi ke rating terkecil (descending) menggunakan insertion sort.

Pada func cetak5Terbaru untuk mencetak urutan 5 buku dengan rating tertinggi. Variabel limit untuk mengantisipasi jika jumlah buku yang diinput kurang dari 5, maka nanti akan mencetak urutan buku sesuai yang diinput.

Pada func cariBuku untuk mencari buku berdasarkan rating menggunakan binary search. Dengan kn dan kr jadi batas pencarian. Setiap iterasi akan mencari indeks tengah, lalu membandingkan rating yang ingin dicari.

Pada func main untuk menjalankan dan memanggil seluruh fungsi secara berurutan, yaitu daftarkanBuku untuk menginput data buku, cetakTerfavorit untuk mencetak buku dengan rating tertinggi, urutBuku untuk mengurutkan buku, cetak5Terbaru untuk mencetak urutan buku dari ratin tertinggi ke rating terkecil, dan cariBuku untuk mencari buku dengan rating yang diinput.