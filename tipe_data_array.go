package main

import "fmt"

func main() {
	// Deklarasi
	var fruits [4]string

	// Inisialisasi
	fruits[0] = "Mango"
	fruits[1] = "Orange"
	fruits[2] = "Melon"
	fruits[3] = "Watermelon"

	var names = [3]string{
		"Dimas", "Setyawan", "Ramadhansyah",
	}

	var numbers = [...]int{ // jumlah data menyesuaikan isi yang diisikan
		1, 2, 3,
	}

	// Akses data
	fmt.Println(names[0])
	fmt.Println(numbers[0])

	// Debugging
	fmt.Println(fruits) // Menghasilkan output dalam bentuk string

	// Array multi-dimensi
	var numbers1 = [2][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}

	// Cara lain
	var numbers2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// Deklarasi make()
	var month = make([]string, 2)

	month[0] = "January"
	month[1] = "Feburary"

	fmt.Println(month)

}

/*
* array: Kumpulan data beritpe sama
*   Kapasitasnya ditentukan dan fix
* Default nilai tiap elemen array sesuai tipe data
* Fungsi:
*   len(): Hitung jumlah elemen pada array
*   make(): Deklarasi variable */
