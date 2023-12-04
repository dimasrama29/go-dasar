package main

import "fmt"

func main() {
	var name string = "Dimas"
	var fruits = [...]int{1, 2, 3}
	var person = map[string]string{
		"name": "Dimas",
		"age":  "24",
	}

	// String
	for i, v := range name {
		fmt.Println(i, v)
	}

	// Slice
	for _, v := range fruits {
		fmt.Println(v)
	}

	// Map
	for i, v := range person {
		fmt.Println(i, v)
	}
}

/*
 * For-range: looping untuk data gabungan (string, array, slice, map)
 */
