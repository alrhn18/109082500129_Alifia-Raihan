package main

import "fmt"

const nMax = 1000

type arrInt [nMax]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func checkJarak(T *arrInt, n int) string {
	var jarak, k int
	var tetap bool
	jarak = T[1] - T[0]
	tetap = true
	k = 2
	for k < n && tetap {
		if T[k]-T[k-1] != jarak {
			tetap = false
		}
		k = k + 1
	}
	if tetap {
		return fmt.Sprintf("Data berjarak %d", jarak)
	}
	return "Data berjarak tidak tetap"
}

func main() {
	var arr arrInt
	var n, x, k int

	fmt.Scan(&x)
	for x >= 0 {
		arr[n] = x
		n = n + 1
		fmt.Scan(&x)
	}

	insertionSort(&arr, n)

	for k < n {
		if k > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[k])
		k = k + 1
	}
	fmt.Println()
	fmt.Println(checkJarak(&arr, n))
}