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