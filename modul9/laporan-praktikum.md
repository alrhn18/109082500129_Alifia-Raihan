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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal1.jpeg)

##### Penjelasan
code diatas untuk membuat deret fibonacci. 

func fibonacci adalah fungsi rekrusif untuk menghitung nilai fibonacci dari ke n. kalo n <= 1 maka akan kembali ke n. jika n > 1 maka akan menggunakan rumus fibonacci(n) = fibonacci(n-1) + fibonacci(n-2).

di func main untuk menginput angka yang mau dibuat jadi deret fibonacci. menggunkan perulangan "for". setiap perulangan akan menghitung fibonaci ke i. hasil dari fibonacci akan di simpan di "hasil". deretan fibonacci akan ditampilkan.


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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal2.jpeg)

##### Penjelasan
code diatas untuk menampilkan pola bintang. 

func bintang merupakan fungsi rekrusif untuk mencetak jumlah bintang sesuai banyaknya yang diinput. jika n == 0 maka program akan berhenti. tetapi jika n > 0 akan mencetak bintang (*) lalu memanggil fungsi bintang(n-1)

func main untuk menginput a, yang merupakan jumlah baris pola. terjadi perulangan dari 1 sampai a, pada iterasi fungsi bintang(i) akan dipanggil lagi untuk untuk mencetak i bintang. fmt.Println() untuk pindah ke baris yang baru.


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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal3.png)

##### Penjelasan
code diatas untuk mencari faktor dari n.

func faktor untuk mencari dan menampilkan faktor faktor dari bilangan a. jika pembagi > a maka akan berhenti. jika a % pembagi == 0 maka pembagi merupakan faktor dari a, maka akan ditampilkan. lalu memanggil fungsi pembagi + 1 untuk mengecek pembagi berikutnya.

func main akan menginput bilangan n yang akan dicari faktornya. lalu memanggil fungsi faktor(n, 1), yang artinya n adalah a dan pembagi adalah 1. lalu program akan mencari faktor dari 1 hingga n.


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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul9/output/output-soal4.png)

##### Penjelasan
code diatas untuk mencari faktor dari n.

func faktor untuk mencari dan menampilkan faktor faktor dari bilangan a. jika pembagi > a maka akan berhenti. jika a % pembagi == 0 maka pembagi merupakan faktor dari a, maka akan ditampilkan. lalu memanggil fungsi pembagi + 1 untuk mengecek pembagi berikutnya.

func main akan menginput bilangan n yang akan dicari faktornya. lalu memanggil fungsi faktor(n, 1), yang artinya n adalah a dan pembagi adalah 1. lalu program akan mencari faktor dari 1 hingga n.