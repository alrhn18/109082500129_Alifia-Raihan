# <h1 align="center">Laporan Praktikum Modul 14 </h1>
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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14/output/output-soal1.jpg)

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14/output/output-soal2.jpg)

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
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul14/output/output-soal3.jpg)

##### Penjelasan
Kode tersebut untuk mencari nilai median dari kumpulan data yang sudah terbaca.

Func median untuk mengurutkan nilai yang diinput dari yang terkecil ke terbesar dengan cara sorting. Bekerja dengan cara mengambil satu elemen a, lalu menggeser elemen elemen sebelumya yang lebih besar ke kanan hingga menjadi terurut dari terkecil ke terbesar.

Func main untuk membaca angka yang diinput. Bilangan selain 0 dan -5313 akan disimpan ke array skor. Saat program membaca 0, func median akan dipanggil lalu akan menghitung median. Jika jumlah data ganjil, maka akan diampil elemen yang ditengah, jika jumlah data genap, maka diambil rata rata dua elemen tengah yang dibulatkan kebawah.

Progran ini akan berhenti ketika program membaca -5313