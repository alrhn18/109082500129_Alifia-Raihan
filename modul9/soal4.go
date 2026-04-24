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