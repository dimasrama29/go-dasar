package main

import "fmt"

func main() {
	var number int = 4
	fmt.Println("Before: ", number)
	fmt.Println(&number)

	change(number, 10)
	fmt.Println("After: ", number)
	fmt.Println(&number)
}

func change(original int, value int) {
	original = value
	fmt.Println(original)
}

/*
* Mengirim data ke function harus menggunakan pointer
*   Karena tidak dalam pointer yang sama*/
