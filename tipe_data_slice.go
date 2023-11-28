package main

import "fmt"

func main() {
	// Deklarasi
	var fruits []string
	fruits = []string{}

	// Inisialisasi
	// fruits = []string{"Mango", "Apple", "Lemon"}

	// Akses data
	fmt.Println(fruits)

	var number = [3]int{1, 2, 3}
	var newNumber = number[0:2] // Slice dari manipulasi array

	fmt.Println(newNumber)

	// Deklarasi make()
	var days = make([]int, 4, 5)

	fmt.Println(len(days))
	fmt.Println(cap(days))
}

/*
* Slice: Reference elemen array
*   Slice bisa dibuat
*   Slice bisa dihasilkan dari manipulasi array (2 index)
*   Slice bisa dibuat dari slice lainnya
* Jika akses menggunakan 1 index, nilai yang didapat adalah pass by value, bukan pass by reference
* Fungsi:
*   len(): Hitung jumlah elemen
*   cap(): Hitung jumlah kapasitas
*   append(): Tambah elemen pada slice
*     Jika len == cap, akan membuat referensi baru
*     Jika len < cap, dimasukkan ke referensi yang sama (Jika masih memungkinkan)
*   copy(): copy elemen dari parameter2(src) ke parameter1(dst)
*     Jika src < dst, akan dicopy semua
*     Jika src > dst, dicopy sesuai cap dst
*     Kembalian berupa angka jumlah elemen yang berhasil di copy
*   make([]string, 3, 3): string dengan len 3 dan cap 3
* Akses dengan 3 index: Menentukan juga cap
* make(): Deklarasi variable*/
