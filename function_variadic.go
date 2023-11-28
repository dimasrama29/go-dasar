package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Printf("Rata-rata %.2f\n", avg)

	// Slice sebagai parameter
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate(numbers...)
	fmt.Printf("Rata-rata %.2f\n", avg)

	yourHobbies("Rama", "Renang", "Lari", "Volly")

	var dataHobby = []string{"Nonton", "Baca"}
	yourHobbies("Ira", dataHobby...)

}

func calculate(numbers ...int) float64 {
	var total int
	for _, v := range numbers {
		total += v
	}
	var avg float64 = float64(total) / float64(len(numbers)) // Melakukan casting ke float64
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesString string = strings.Join(hobbies, ",")
	fmt.Printf("Nama\t: %s\nHobbies\t: %s\n", name, hobbiesString)
}

/*
* Function Variadic: Fungsi dengan parameter sejenis tanpa batas
*   Sifat mirip seperti slice
* Cara slice sebagai parameter
*   Tambahan ... setelah variable
* Bisa dikombinasikan dengan parameter biasa, tetapi variadic harus di parameter akhir*/
