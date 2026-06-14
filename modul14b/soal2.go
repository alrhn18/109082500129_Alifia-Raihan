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