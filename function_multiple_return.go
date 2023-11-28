package main

import (
	"fmt"
	"math"
)

func main() {
	var getKeliling, getLuas = calculated(70)
	fmt.Printf("Keliling\t: %.2f\n", getKeliling)
	fmt.Printf("Luas\t\t: %.2f\n", getLuas)
}

func calculated(d float64) (float64, float64) {
	// Keliling
	var keliling float64 = math.Pi * d

	// Luas
	var luas float64 = math.Pi * math.Pow(d/2, 2)

	return keliling, luas
}
