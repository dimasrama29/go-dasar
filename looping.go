package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Hello-%d \n", i)
	}

	// Hanya kondisi
	var i = 0
	for i > 5 {
		fmt.Println("Angka", i)
		i++
	}

	// Tanpa argumen
	for {
		if i > 10 {
			fmt.Println(i)
		}
	}

}

/*
 * Loop: Mengulang-ulang blok kode sesuai dengan kondisi acuan yang terpenuhi
 * Jika for tanpa kondisi akan sama dengan for true (Tanpa henti)
 */
