package main

import (
	"fmt"
	"math"
)

func main() {
	var getKeliling, getLuas = calculated(70)
	fmt.Printf("Keliling \t: %.2f\n", getKeliling)
	fmt.Printf("Luas \t\t: %.2f\n", getLuas)
}

func calculated(d float64) (keliling float64, luas float64) {
	keliling = math.Pi * d
	luas = math.Pi * math.Pow(d/2, 2)

	return
}

/*
* Tidak perlu deklarasi variable kembali di dalam blok
* Karena sudah definisikan di return, tidak perlu deklarasi variable kembali*/
