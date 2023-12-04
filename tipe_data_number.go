package main

import "fmt"

func main() {
	var positiveNumber int8 = 100
	var negativeNumber int = -98765

	fmt.Printf("%d dan %d \n", positiveNumber, negativeNumber)

	var decimalNumber float64 = 1.23456789

	fmt.Printf("%f dan %.3f \n", decimalNumber, decimalNumber)
}

/*
 * Non-desimal:
 * 	Bilangan cacah
 *  	uint8, uint16, uint32, uint64
 *  Bilangan bulat
 * 		int8, int16, int32, int64
 * Desimal:
 * 	float32, float64
 * Alias:
 * 	uint: uint32 / uint64 (Tergantung nilai)
 *  byte: uint8
 *  rune: int32
 *  int: int32 / int64 (Tergantung nilai)
 * Pilih tipe data sesuai jangkauan, efeknya ke alokasi memori variable (pemakaian lebih optimal)
 * Default value: 0
 */
