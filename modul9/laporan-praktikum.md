# <h1 align="center">Laporan Praktikum Modul 9 </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Mencari Titik di Dalam Lingkaran 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.

#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))
}

func dalam(cx, cy, x, y, r float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y         float64
	)

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dlm1 := dalam(cx1, cy1, x, y, r1)
	dlm2 := dalam(cx2, cy2, x, y, r2)

	if dlm1 && dlm2 {
		fmt.Println("titik dalam lingkaran 1 dan 2")
	} else if dlm1 {
		fmt.Println("titik dalam lingkaran 1")
	} else if dlm2 {
		fmt.Println("titik dalam lingkaran 2")
	} else {
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal1.jpg)

##### Penjelasan
codingan diatas untuk mencari posisi titik lingkaran, dengan:

cx, cy adalah titik tengah lingkaran,
x, y adalah titik yang akan di cek, dan
r adalah jarak maksimal dari pusat

func jarak untuk mencari jarak antara titik tengah (cx, cy) dan titik (x,y).


func dalam untuk mencari kebenaran dari jarak yang sudak di cari di func jarak dengan jarak maksimal.

jika d <= r maka titik tersebut berada di dalam lingkaran,
jika d > r maka titik tersebut berada di luar lingkaran.


pada bagian func main menjelaskan:

dlm1 untuk mengitung jarak antara titik pusat(cx1, cy1) dan titik(x, y), dan menjelaskan kebenaran antara jarak yang sudah dihitung dengan jarak maksimal(r1)
dlm2 untuk menghitung jarak antara titik pusat(cx2, cy2) dan titik(x, y), dan menjelaskan kebenaran antara jarak yang sudah dihitung dengan jarak maksimal(r2)

pada baris ke 120-121 itu menjelaskan kebenarannya, antara TRUE atau FALSE, maka outputnya:
jika TRUE dan TRUE maka outputnya "titik dalam lingkaran 1 dan 2",
jika yang TRUE hanya di "dlm1" maka outputnya "titik di dalam lingkaran 1",
jika yang TRUE hanya di "dlm2" maka outputnya "titik di dalam lingkaran 2", dan
jika FALSE semua maka outputnya "titik di luar lingkaran 1 dan 2".


## Array

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut: keseluruhan isi array, elemen elemen array dengan indeks ganjil dan genap, elemen elemen array dengan kelipatan bilangan x, menghapus elemen array pada indeks tertentu, rata rata dari bilangan dalam array, simpangan baku yang ada di dalam array, dan frekuensi dari suatu bilangan.

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal2.jpg)

##### Penjelasan
code diatas untuk menampilkan keseluruhan isi array, elemen elemen array dengan indeks ganjil dan genap, elemen elemen array dengan kelipatan bilangan x, menghapus elemen array pada indeks tertentu, rata rata dari bilangan dalam array, simpangan baku yang ada di dalam array, dan frekuensi dari suatu bilangan. 

pada baris 105, akan membuat array sebanyak n, dan akan mengisi elemen satu satu. pada baris 115 sampai 130, akan membuat putput elemen dari index ganjil dan genap, dengan cara di modulo 2. pada baris 132, akan membuat index kelipatan x. pada baris 143, akan memnghapus elemen dari index tertentu. pada baris 150, akan mencari rata rata dari elemen tersebut, dengan cara menambahkan semua elemen lalu dibagi banyaknya elemen, lalu diubah ke dalam float64 agar hasilnya desimal. pada baris 158, untuk mencari simpangan baku. pada baris 166, akan mencari frekuensi yang muncul di array, dengan cara loop seluruh array dan menghitung berapa kali nilai yang diinput muncul.


## Pertandingan Bola

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.

#### soal3.go

```go
package main 

import "fmt"

func main () {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [] string
	
	fmt.Print("klub a: ")
	fmt.Scan(&klubA)

	fmt.Print("klub b: ")
	fmt.Scan(&klubB)

	pertandingan := 1
	skorA, skorB = 0, 0
	for skorA >= 0 && skorB >= 0{
		fmt.Printf("pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA >= 0 && skorB >= 0 {
			if skorA > skorB{
				hasil = append(hasil, klubA)
			}else if skorA < skorB{
				hasil = append(hasil, klubB)
			}else if skorA == skorB{
				hasil = append(hasil, "draw")
			}
		}

		pertandingan++
	}

	fmt.Println(" ")
	for i := 0; i < len(hasil); i++{
		fmt.Printf("hasil %d : %s", i+1, hasil[i])
		fmt.Println(" ")
	}

	fmt.Println("pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal3.jpg)

##### Penjelasan
code diatas untuk mencari pemenang dalam pertandingan bola antara MU dan Inter.

pada baris 206-210 untuk menginput nama klub bola. pada baris 212-229 untuk menginput skor pertandingan, dan menentukan siapa pemenang dari setiap pertandingan, dan memberi output skor per-pertandingan. pada baris 231-237 untuk memberi output siapa pemenang dari per-pertandingan.


## Palindrom

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isi_array(t *tabel, n *int){
	var input string
	fmt.Scan(&input )

	*n = 0
	for i := 0; i < len(input) && *n < int(NMAX); i++{
		t[*n] = rune(input[i])
		*n++
	}
}

func cetak_array(t *tabel, n int){
	for i := 0; i < n; i++{
		fmt.Printf("%c", t[i])
	}
	fmt.Println(" ")
}

func balikan_array(t *tabel, n int){
	i := 0
	j := n -1

	for i < j{
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func palindrom(t *tabel, n int)bool{
	i := 0
	j := n -1

	for i < j {
		if t[i] != t[j]{
			return false
		}
		i++
		j--
	}

	return true
}

func main () {
	var tab tabel
	var m int

	fmt.Print("kalimat: ")
	isi_array(&tab, &m)

	palindrom := palindrom(&tab, m)

	balikan_array(&tab, m)

	fmt.Print("membalikkan kalimat: ")
	cetak_array(&tab, m)

	fmt.Print("palindrom: ", palindrom)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal4.jpg)

##### Penjelasan
code diatas untuk mencari palindrom dan output dari kebalikan kalimat yang kita input

pada func isi_array untuk membaca inputan kalimat. pada func cetak_array untuk mencetak setiap karater dalam array. pada func balikan_array untuk memberi output kebalikan kalimat yang diinput diawal. pada func palindrom untuk mengecek apakah kalimat yang diinput merupakan palindrom atau tidak. pada func main untuk pengecekan palindrom sebelum array dibalik, dan untuk menginput kalimat, dan memberi output dari kalimat yang sudah dibalik, dan true atau false nya palindrom.