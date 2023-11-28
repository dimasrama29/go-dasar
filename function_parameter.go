package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{
		"Dimas", "Setyawan", "Ramadhansyah",
	}
	printMessage("Hello", data)

	divideNumber(10, 0)
}

func printMessage(message string, data []string) {
	var dataString string = strings.Join(data, " ") // Menggabungkan dengan pembatas spasi
	fmt.Printf("%s %v\n", message, dataString)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Penyebut tidak bisa 0\n")
		return
	} else {
		var result int = m / n
		fmt.Printf("%d / %d = %d\n", m, n, result)
	}
}

/*
* Parameter: Variable yang dimasukkan saat pemanggilan fungsi
* Jika seluruh parameter memiliki tipe data yang sama, bisa ditulis sekali di akhir*/
