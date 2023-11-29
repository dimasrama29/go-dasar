package main

import "fmt"

func main() {
	// Deklarasi
	var numberB *int

	// Inisialisasi
	var numberA int = 4
	numberB = &numberA

	fmt.Println("Value:", numberA)
	fmt.Println("Address:", &numberA)

	fmt.Println("Value:", *numberB)
	fmt.Println("Address:", numberB)

	fmt.Println("")

	fmt.Println("Before:", numberA)

	// Perubahan nilai
	*numberB = 5
	fmt.Println("After:", *numberB)

	// Deklarasi
	var name *string = new(string)
	*name = "Rama"

	fmt.Println(name)
	fmt.Println(*name)
}

/*
* Pointer: Reference / alamat memori
* Variable pointer: Variable yang berisi alamat memori suatu nilai
* Variable di go pass by value
* 	Jika kirim variable ke function, method, atau variable lain yang dikirim adalah hasil duplikasinya
*   Variable yang reference / memiliki alamt memori yang sama, saling berhubungan satu sama lain
* & ampersand (Referencing): Mengambil nilai pointer dari variable biasa
* * asterisk (Deferencing): Mengambil nilai asli dari variable pointer
 * New: Membuat variable pointer dengan tipe data tertentu
 * Nilai default: nil
 *  Nilai default sesuai dengan tipe datanya*/
