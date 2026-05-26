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