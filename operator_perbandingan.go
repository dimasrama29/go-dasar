package main

import "fmt"

func main() {
	var value int = 9 + 10/2 - 3
	var isEqual bool = value == 15

	fmt.Println(isEqual)
}

/*
* Operator perbandingan:
*   ==, >, <, >=, <=, !
* Hasilnya berupa boolean (true / false)*/
