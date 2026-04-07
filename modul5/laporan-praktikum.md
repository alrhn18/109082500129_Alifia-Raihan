# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[nama] - [NIM]</p>

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul5/output/output-soal1.jpg)

##### Penjelasan
[penjelasan]


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
[penjelasan]


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
[penjelasan]