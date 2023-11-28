package main

import "fmt"

func main() {
	// Deklarasi | Manifest typing
	var firstName string

	// Inisialisasi
	firstName = "Dimas"

	// Type inference
	var middleName = "Setyawan"
	lastName := "Ramadhansyah"

	fmt.Printf("Hello %s %s! \n", firstName, lastName)

	// Deklarasi multi-variable
	var first, second, third int

	// Inisialisasi
	first, second, third = 1, 2, 3
	forth, fifth, sixth := 5, "6", 7 // Jika menggunakan type inference, bisa berbeda tipe data
}

/*
* 2 cara penulisan
*   Manifest Typing: Menuliskan tipe data
*   Type inference: Tanpa menuliskan tipe data
*     Tipe data menyesuaikan nilainya
*     := hanya untuk assignment
*     := hanya digunakan di dalam block fungsi
* Semua variable yang dideklarasikan harus di tampilkan*/
