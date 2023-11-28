package main

import "fmt"

func main() {
	// Deklarasi
	var person map[string]int
	person = map[string]int{}

	// Inisialisasi
	person["January"] = 1
	person["February"] = 2

	var fruits = map[string]string{
		"Name": "Mango",
		"ID":   "20",
	}

	// Akses
	fmt.Println(person["January"])
	fmt.Println(person["May"]) // Jika belum di set, akan mengembalikan nilai default tipe data

	// Deklarasi make()
	var number1 = make(map[string]string)

	fmt.Println(number1)

	// Cek item tertentu
	var value, isExist = fruits["Name"] // variable kedua mengembalikan boolean

	if isExist == true {
		fmt.Println(value)
	} else {
		fmt.Println("Doesn't Exist")
	}
}

// Kombinasi map dan slice
var person1 = []map[string]string{
	map[string]string{"Name": "Dimas", "Age": "24"},
	map[string]string{"Name": "Ira", "Age": "21"},
}

// Atau
var person2 = []map[string]string{
	{"Name": "Dimas", "Age": "24"},
	{"Address": "Kalasan"},
}

/*
* Map: Tipe data assosiatif yang berbentuk key-value
*   Default nilai variable map adalah nil
* fungsi:
*   make(): Deklarasi variable
*   delete(): Menghapus item tertentu
*   len(): Hitung jumlah item*/
