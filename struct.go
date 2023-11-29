package main

import "fmt"

// Deklarasi struct
type Student struct {
	Name string
	Age  int
}

func main() {
	// Buat variable object / object struct
	var s1 Student

	// Set nilai property variable object
	s1.Name = "Dimas Setyawan Ramadhansyah"
	s1.Age = 24

	fmt.Println("Name:", s1.Name)
	fmt.Println("Age: ", s1.Age)

	// Inisialisasi
	// cara 1
	var s2 = Student{}
	s2.Name = "Ira"
	s2.Age = 21

	// cara 2
	var s3 = Student{
		Name: "Kris",
		Age:  51,
	}

	// cara 3
	var s4 = Student{"Rama", 24}

	fmt.Println(s2.Name)
	fmt.Println(s3.Name)
	fmt.Println(s4.Name)

}

/*
* Struct: Kumpulan definisi variable / property dan fungsi / method sebagai tipe data baru
*   Nilai awal property sesuai tipe data
* */
