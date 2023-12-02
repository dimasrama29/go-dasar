package main

import "fmt"

func main() {
	var message string = "Hello"

	fmt.Printf("Message: %s \n", message)

	// Backtick
	var pesan string = `Nama saya adalah "Dimas 
  Setyawan Ramadhansyah" \n`

	fmt.Println(pesan)
}

/*
 * Backtick: Semua karakter di dalamnya tidak di escape
 *   Termasuk "", '', newline
 *   Default nilai: ""
 */
