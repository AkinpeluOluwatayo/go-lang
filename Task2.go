package assignment

import "fmt"

func DrawTriangle() {
	var baselength int

	fmt.Print("Enter the base length of the triangle (1-10): ")
	fmt.Scan(&baselength)

	if baselength < 1 || baselength > 10 {
		fmt.Println("Please enter a number between 1 and 10.")
		return
	}

	fmt.Println("Generating Triangle:")

	for index := 1; index <= baselength; index++ {
		for inner := 1; inner <= index; inner++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
