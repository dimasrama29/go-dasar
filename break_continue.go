package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			continue
		}

		if i >= 50 {
			break
		}

		fmt.Println(i)
	}
}

/*
 * Break: Menghentikan perulangan secara paksa
 * Continue: Memaksa maju ke perulganan berikutnya
 */
