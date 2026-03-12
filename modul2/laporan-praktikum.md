# <h1 align="center">Laporan Praktikum Modul 2 </h1>
<p align="center">Alifia Raihan Nugraheni - 109082500129</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

#### soal1.go

```go
package main 

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)

}	

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul2/output/output-soal1.jpg)

[penjelasan]
codingan diatas untuk menukar isi variabel. dengan 4 string, yang dimana a, b, c untuk menginput string dan temp untuk variabel sementara sebelum diubah.

"temp = satu" artinya temp adalah nilai a yang merupakan variabel pertama yang disimpan terlebih dahulu.
"satu = dua" artinya untuk output yang pertama keluar adalah b yang merupakan variabel kedua
"dua = tiga" artinya untuk output kedua yang keluar adalah c yang merupakan variabel ketiga
"tiga = temp" artinya untuk output ketiga yang keluar adalah a yang merupakan variabel pertama yang disimpen diawal

pada output awal masih normal tidak ada perubahan masih sesuai dengan apa yang diinput (a, b, c).
pada output akhir output akan berubah yang tadinya a menjadi b, yang tadinya b jadi c, dan yang tadinya c jadi a (b, c, a)

## Unguided

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

#### soal2.go

```go
package main 

import "fmt"

func main () {
	var (
		a, b, c, d string
		i int
		e bool
	)

	e = true
	for i = 1; i <= 5; i++ {
		fmt.Printf("percobaan ke-%d: ", i)
		fmt.Scan(&a, &b, &c, &d)

		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu"{
			e = false
		}
	}
	fmt.Println("hasil:", e)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul2/output/output-soal2.jpg)

[penjelasan]
codingan diatas merupakan hasil dari praktikum kimia kelas ipa. kelas ipa melakukan sebanyak 5 kali percobaan. jika semua percabaan itu hasilnya merah, kuning, hijau, dan ungu maka outputnya akan "TRUE" namun apabila ada salah satu percobaan yang hasilnya tidak merah, kuning, hijau, atau ungu maka outputnya akan "FALSE"

## Unguided

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### soal3.go

```go
package main

import "fmt"

func main() {
	var a int

	fmt.Print("berat parsel (gram): ")
	fmt.Scan(&a)

	b := a / 1000
	c := a % 1000
	d := b * 10000

	fmt.Printf("detail berat: %v kg + %v gr", b, c)
	fmt.Println(" ")

	if b > 10 {
		e := 0
		fmt.Printf("detail harga: Rp.%v + Rp. %v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	} else if b < 10 && c >= 500 {
		e := c * 5
		fmt.Printf("detail harga: Rp.%v + Rp.%v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	} else {
		e := c * 15
		fmt.Printf("detail harga: Rp.%v + Rp.%v", d, e)
		fmt.Println(" ")
		fmt.Println("total biaya: Rp.", d+e)
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alrhn18/109082500129_Alifia-Raihan/blob/main/modul2/output/output-soal3.jpg)

[penjelasan]
codingan diatas untuk menghitung biaya jasa pengiriman menggunakan berat parsel. berat parsel yang awalnya dalam bentuk gram di ubah menjadi kg dan sisanya dalam gram. dengan biaya pengiriman Rp. 10000 per kg, jika sisa berat >500 gram dikenakan biaya tambahan Rp. 5 per gram, jika sisa berat <500 gram dikenakan biaya tambahan Rp. 15 per gram, dan jika pengiriman berat parsel >10 kg tidak dikenakan biaya tambahan sisa beratnya. 

pada baris 108 untuk mencari berat parsel dalam bentuk kg.
pada baris 109 untuk mencari berat sisa parsel.
pada baris 110 untuk menghitung harga berat parsel per kg.

pada baris 117-119 untuk menghitung biaya pengiriman parsel dengan berat lebih dari 10 kg.
pada baris 121-124 untuk menghitung biaya pengiriman parsel dengan berat kurang dari 10 kg dan sisa berat lebih dari 500 gram.
pada baris 126-129 untuk menghitung biaya pengiriman parsel dengan berat kurang dari 10 kg dan sisa berat kurang dari 500 gram.
