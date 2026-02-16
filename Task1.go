package assignment

import "fmt"

func ValueOfX() {
	x := 7
	y := 3

	y++
	x = y
	fmt.Println("x is ", x)
}
