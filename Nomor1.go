package main

import "fmt"

func main() {
	var i, n, a, b, positif, negatif int // i = interaksi; n = banyak tempat; a = energi pertama; b = energi kedua ; positif = menampung sluruh energi positif; negatif = seluruh energi negatif

	positif = 0
	negatif = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&a, &b)
		if a > 0 {
			positif += a
		} else {
			negatif += a
		}
		if b > 0 {
			positif += b
		} else {
			negatif += b
		}
	}
	fmt.Println("Negative:", negatif)
	fmt.Println("positive:", positif)
}
