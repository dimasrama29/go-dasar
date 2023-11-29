package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Dimas", "Setyawan", "Ramadhansyah"}
	var dataContains = filter(data, func(v string) bool {
		return strings.Contains(v, "h")
	})

	var dataLength = filter(data, func(v string) bool {
		return len(v) == 5
	})
	fmt.Println(dataContains)
	fmt.Println(dataLength)
}

type Filter func(string) bool

func filter(data []string, callback Filter) []string {
	var result []string
	for _, v := range data {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

/*
* Fungsi bisa dijadikan tipe data variable, Sehingga bisa dijadikan parameter
* TEMUAN
*	Bisa digunakan untuk lebih dari satu func dengan tipe data sama
* */
