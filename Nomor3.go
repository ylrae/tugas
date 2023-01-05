package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	konversi(x, &y)
	fmt.Print(y)
}

func pangkat(a, b int) int {
	var i, hasil int
	hasil = 1
	for i = 1; i <= b; i++ {
		hasil *= a
	}
	return hasil
}

func konversi(biner int, desimal *int) {
	var x, j int
	j = 0
	for biner > 0 {
		x = biner % 10
		*desimal += x * pangkat(2, j)
		j++
		biner /= 10
	}
}
