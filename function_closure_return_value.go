package main

import "fmt"

func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var howMany, getNumbers = findMax(max, numbers)
	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", getNumbers()) // Memanggil fungsi karena return adalah fungsi
}

func findMax(max int, data []int) (int, func() []int) {
	var r []int
	for _, v := range data {
		if v <= max {
			r = append(r, v)
		} else {
			continue
		}
		/*
			Atau return fungsi dimasukkan ke variable dahulu
			var getNumbers = func() []int {
				return r
			}
			return len(r), getNumbers
		*/
	}
	return len(r), func() []int {
		return r
	}
}
