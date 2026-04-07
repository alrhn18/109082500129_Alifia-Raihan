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