package main

import "fmt"

func main() {
	var nilai int = 8

	if nilai > 9 {
		fmt.Println("Sempurna")
	} else if nilai > 8 {
		fmt.Println("Bagus")
	} else {
		fmt.Println("Belajar lagi")
	}

	// Variable Temporary: Hanya bisa digunakan di blok ditempatkan
	var point float32 = 0.84

	if percent := point * 100; percent > 80 {
		fmt.Println("A")
	} else if percent > 70 {
		fmt.Println("B")
	} else {
		fmt.Println("Coba lagi")
	}

}

/*
* Conditional: Mengontrol alur program
* Ada 2 macam Conditional:
*   If else
*   Switch
* Variable temporary: Hanya bisa digunakan di blok ditempatkan
*   Hanya bisa menggunakan := */
