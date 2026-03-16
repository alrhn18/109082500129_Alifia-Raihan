# <h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">Alifia Raihan Nugraheni - 109082500129</p>

## Faktorial dan Kombinasi 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas?

#### soal1.go

```go
package main 

import "fmt"

func faktorial(n int)int {
	hasil := 1
	for i := 1; i <= n; i++{
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int)int{
	return faktorial(n) / faktorial(n-r)	
}

func combinasi(n, r int)int{
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main () {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a,c), combinasi(a,c))
	fmt.Println(permutasi(b,d), combinasi(b,d))
}	

```


##### Output 
(https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul3/output/ouput-soal1.jpg)

##### Penjelasan

codingan diatas untuk mencari permutasi dan kombinasi, dengan:

func faktorial untuk mencari faktorial yang dicari. 
func permutasi untuk mencari permutasi. 
func combinasi untuk mencari coombinasi.

func main, menjelaskan:
untuk println pertama untuk output permutilasi dan kombinasi dari inputan a dan c, dan
untuk println kedua untuk output permutilasi dan kombinasi dari inputan b dan d.

## Fungsi

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

#### soal2.go

```go
package main 

import "fmt"

func f (x int) int{
	return x * x
}

func g (x int) int{
	return x - 2
}

func h (x int) int{
	return x + 1
}

func main () {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

```


##### Output 
(https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul3/output/output-soal2.jpg)

##### Penjelasan

codingan diatas mencari fungsi komposisi (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c), dengan:
a, b, dan c adalah angka yang diinput.

func f, func g, dan func h untuk menghitung angka yang diinput dengan menggunakan soal yang sudah tertera.

pada bagian func main menjelaskan:
untuk println pertama untuk mencari (fogoh)(a),
untuk println kedua untuk mencari (gohof)(b), dan
untuk println ketiga untuk mencari(hofog)(c).

## Lingkaran

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut.

#### soal3.go

```go
package main

import "fmt"
import "math"

func jarak (a, b, c, d float64)float64{
	return math.Sqrt(((a-c)*(a-c)) + ((b-d)*(b-d)))
}

func dalam (cx, cy, x, y, r float64)bool{
	return jarak (cx, cy, x, y) <= r
}

func main () {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
	)

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dlm1 := dalam(cx1, cy1, x, y, r1)
	dlm2 := dalam(cx2, cy2, x, y, r2)

	if dlm1 && dlm2 {
		fmt.Println("titik dalam lingkaran 1 dan 2")
	}else if dlm1 {
		fmt.Println("titik dalam lingkaran 1")
	}else if dlm2 {
		fmt.Println("titik dalam lingkaran 2")
	}else{
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}

```


##### Output 
(https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul3/output/output-soal3.jpg)

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