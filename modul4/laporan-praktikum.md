# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center"> Alifia Raihan Nugraheni - 109082500129 </p>

## Permutasi dan Kombinasi 

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

func P(n, r int)int{
	return faktorial(n) / faktorial(n-r)
}

func C(n, r int)int{
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main () {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(P(a, c), C(a, c))
	fmt.Println(P(b, d), C(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul4/output/output-soal1.jpg)

##### Penjelasan
code diatas untuk mencari permutasi dan kombinasi. 

func faktorial untuk menghitung nilai faktorial dari bilangan n, dengan cara perulangan menggunakan for. di for nanti akan melakukan perulangan dari 1 sampai n, dan hasil merupakan hasil perkalian dari 1 sampai n.

func P untuk menghitung permutasi, dengan rumus P(n, r) = n! / (n - r)!. cara menghitung permutasi memanfaatkan func faktorial, yang sudah di buat.

func C untuk menghitung kombinasi, dengan rumus C(n, r) = n! / (r! * (n - r)!). cara menghitung kombinasi memanfaatkan func faktorial, yang sudah di buat.

func main digunakan untuk menginput variabel yang bersifat int, yaitu a, b, c, d. lalu akan menampilkan hasil permutasi dan kombinasi (a, c) dan (b, d) menggunakan fungsi yang udah di buat sebelumnya. 


## Menghitung Skor Kompetisi Pemrograman Tingkat Nasional

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.

#### soal2.go

```go
package main
import "fmt"

// prosedur hitung skor
func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)

		if waktu < 300 {
			*soal = *soal + 1
			*skor = *skor + waktu
		}
	}
}

func main() {
	var nama string
	var pemenang string
	var soal, skor int
	var maxSoal, minSkor int
	var pertama bool

	fmt.Scan(&nama)

	pertama = true

	for nama != "Selesai" {

		hitungSkor(&soal, &skor)

		if pertama {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
			pertama = false
		} else {
			if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
				pemenang = nama
				maxSoal = soal
				minSkor = skor
			}
		}

		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul4/output/output-soal2.jpeg)

##### Penjelasan
code diatas untuk mencari pemenang kompetisi pemrograman, dari jumlah soal yang bisa dikerjakan dan lama waktu mengerjakan soal.

func hitungSkor untuk mengihtung jumlah soal yang dapat di kerjakan oleh peserta, dan lama waktu yang dibutuhkan peserta. jika waktu pengerjaan lebih dari 300 menit maka tidak masuk ke penjumlahan soal, tetapi tetap bisa menginput waktu pengerjaan selama belum sampai 8 kali. dengan *soal dan *skor diinisialisasi sebagai 0

func main untuk menginput: nama untuk menginput angka, pertama untuk sebagain penyimpanan nama peserta yang menang, soal dan skor untuk meniyimpan total soal dan total skor yang dikerjakan masing masing peserta, maxSoal untuk menyimpan jumlah soal terbanyak yang bisa dikerjakan peserta, minSkor untuk menyimpan waktu tercepat, pertama adalah sebagai penanda untuk peserta pertama yang melakukan kompetisi. 

program akan menginput nama peserta, dan akan berhenti jika menginput "Selesai". di dalam perulangan akan memanggil hitungSkor untuk menghitung jumlah soal dan lama waktu pengerjaan. di dalam perulangan akan terjadi pembandingan jumlah skor dan lama pengerjaan. jika sudah menginput "Selesai" maka akan memberi output nama pemenanng, jumlah soal yang bisa dikejain oleh peserta, dan lama waktu pengerjaan.