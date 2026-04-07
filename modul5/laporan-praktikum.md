# <h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Deret fibonacci 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var a int

	fmt.Scan(&a)
	
	for i := 0; i <= a; i++ {
		hasil := fibonacci(i)
		fmt.Print(hasil, " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul5/output/output-soal1.jpeg)

##### Penjelasan
code diatas untuk membuat deret fibonacci. 

func fibonacci adalah fungsi rekrusif untuk menghitung nilai fibonacci dari ke n. kalo n <= 1 maka akan kembali ke n. jika n > 1 maka akan menggunakan rumus fibonacci(n) = fibonacci(n-1) + fibonacci(n-2).

di func main untuk menginput angka yang mau dibuat jadi deret fibonacci. menggunkan perulangan "for". setiap perulangan akan menghitung fibonaci ke i. hasil dari fibonacci akan di simpan di "hasil". deretan fibonacci akan ditampilkan.


## Pola Bintang

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul5/output/output-soal2.jpeg)

##### Penjelasan
code diatas untuk menampilkan pola bintang. 

func bintang merupakan fungsi rekrusif untuk mencetak jumlah bintang sesuai banyaknya yang diinput. jika n == 0 maka program akan berhenti. tetapi jika n > 0 akan mencetak bintang (*) lalu memanggil fungsi bintang(n-1)

func main untuk menginput a, yang merupakan jumlah baris pola. terjadi perulangan dari 1 sampai a, pada iterasi fungsi bintang(i) akan dipanggil lagi untuk untuk mencetak i bintang. fmt.Println() untuk pindah ke baris yang baru.


## Faktor Bilangan

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.

#### soal3.go

```go
package main 

import "fmt"

func faktor (a, pembagi int){
	if pembagi > a{
		return
	}

	if a % pembagi == 0{
		fmt.Print(pembagi, " ")
	}

	faktor(a, pembagi + 1)
}

func main () {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul5/output/output-soal2.jpeg)

##### Penjelasan
code diatas untuk mencari faktor dari n.

func faktor untuk mencari dan menampilkan faktor faktor dari bilangan a. jika pembagi > a maka akan berhenti. jika a % pembagi == 0 maka pembagi merupakan faktor dari a, maka akan ditampilkan. lalu memanggil fungsi pembagi + 1 untuk mengecek pembagi berikutnya.

func main akan menginput bilangan n yang akan dicari faktornya. lalu memanggil fungsi faktor(n, 1), yang artinya n adalah a dan pembagi adalah 1. lalu program akan mencari faktor dari 1 hingga n.