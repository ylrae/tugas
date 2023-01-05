package main

import "fmt"

func main() {
	var j, pengunjung, sebelum, penurunan, jumlah int
	penurunan = 0
	sebelum = 0
	j = 0
	jumlah = 0
	for penurunan != 3 {
		j = j + 1
		fmt.Print("Hari ", j, " : ")
		fmt.Scan(&pengunjung)
		for pengunjung < 0 || pengunjung > 100 {
			fmt.Print("Hari ", j, " : ")
			fmt.Scan(&pengunjung)
		}
		jumlah += pengunjung
		if pengunjung < sebelum {
			penurunan += 1
		} else {
			penurunan = 0
		}
		sebelum = pengunjung
	}
	fmt.Println("Museum buka selama ", j, "hari")
	fmt.Printf("rata-rata pengunjung %.2f orang", float64(jumlah)/float64(j))
}
