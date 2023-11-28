package main

import "fmt"

func main() {
	var left bool = false
	var right bool = true

	var leftAndRigth bool = left && right
	fmt.Printf("Right && Left \t %t \n", leftAndRigth)

	var leftOrRigth bool = left || right
	fmt.Printf("Right || Left \t %t \n", leftOrRigth)
}

/*
* Operator logika:
*   &&, ||, !
* Membandingkan boolean dan menghasilkan nilai bool */
