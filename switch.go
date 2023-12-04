package main

import "fmt"

func main() {
	var point int = 9

	switch point {
	case 9:
		fmt.Println("A")
	case 8, 7:
		fmt.Println("B")
	default:
		{ // Digunakan untuk banyak statement
			fmt.Println("Coba lagi")
			fmt.Println("END")
		}
	}

	// Switch gaya if-else
	switch {
	case point > 9:
		fmt.Println("A")
		fallthrough
	case point > 8:
		fmt.Println("B")
	default:
		fmt.Println("You need learn")
	}
}

/*
 * Switch: Fokus pada satu variable
 *   Ketika case terpenuhi, tidak dilanjutkan pengecekan berikutnya
 * Fallthrough: Memaksa pengecekan diteruskan ke satu case selanjutnya
 */
