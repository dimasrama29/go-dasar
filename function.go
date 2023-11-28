package main

import "fmt"

func main() {
	// Fungsi bisa dipanggil berkali-kali
	sayHello()
	sayHello()
	sayHello()
}

func sayHello() {
	fmt.Printf("Hello World \n")
}

/*
* Dengan fungsi kode menjadi lebih modular dan DRY (Don't Repeat Yourself)
*	Modular: Memecah masalah menjadi masalah lebih kecil
*	DRY: Tidak perlu mengulang tulisan yang sama berkali-kali*/
