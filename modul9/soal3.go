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