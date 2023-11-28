package main

import "fmt"

func main() {
  // Deklarasi 
  var numberB *int 

  // Inisialisasi
  var numberA int = 4
  numberB = &numberA

  fmt.Println(numberA)
  fmt.Println(&numberA)

  fmt.Println(*numberB)
  fmt.Println(numberB)

  fmt.Println("")

  // Perubahan nilai 
  *numberB = 5

  fmt.Println(numberA)
  fmt.Println(&numberA)

  fmt.Println(*numberB)
  fmt.Println(numberB)

	// Deklarasi
	var name *string = new(string)
	*name = "Rama"

	fmt.Println(name)
	fmt.Println(*name)
}

/*
* Pointer: Reference / alamat memori
*   Variable yang reference / memiliki alamt memori yang sama, saling berhubungan satu sama lain 
* Variable biasa bisa diambil nilai pointernya dengan menambahkan ampersand (&) / Referencing 
* Variable pointer bisa diambil nilai aslinya dengan menambahkan asterisk (*) / Deferencing 
 * New: Membuat variable pointer dengan tipe data tertentu
 * Nilai default variable pointer: nil 
 *  Nilai default sesuai dengan tipe datanya*/
