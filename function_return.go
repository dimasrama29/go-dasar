package main

import "fmt"

func main() {
	var hello string = getHello("Dimas")
	fmt.Printf("%s\n", hello)
}

func getHello(name string) string {
	var message string = "Hello " + name
	return message
}

/*
* Void: Fungsi tanpa nilai kembalian
* Saat keyword return dieksekusi, semua proses dalam blok fungsi akan berhenti*/
