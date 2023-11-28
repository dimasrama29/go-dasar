package main

import "fmt"

func main() {
	// Deklarasi closure
	var getMinMax func([]int) (int, int)

	// Inisialisasi
	getMinMax = func(numbers []int) (int, int) {
		var min, max int
		for i, v := range numbers {
			if i == 0 {
				min, max = v, v
			} else if v > max {
				max = v
			} else if v < min {
				min = v
			}
		}
		return min, max
	}

	var number = []int{2, 3, 4, 3, 4, 2, 3}

	var min, max int = getMinMax(number)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", number, min, max)

	// IIFE: Langsung dikirim argument diakhir closure
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, v := range numbers {
			if v < min {
				continue
			} else {
				r = append(r, v)
			}
		}
		return r
	}(3) // Memasukkan nilai ke variable min

	fmt.Println("Original number:", numbers)
	fmt.Println("Filtered number:", newNumbers)

}

/*
* Closure: Fungsi yang bisa disimpan di dalam variable
*   Maka bisa membuat:
*     Fungsi di dalam Fungsi
*     Fungsi Mengembalikan fungsi
* Closure = Anonymous function
*   Digunakan untuk membungkus proses yang digunakan hanya sekali
* BELUM SELESAI*/
