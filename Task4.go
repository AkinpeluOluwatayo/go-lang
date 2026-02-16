package assignment

import "fmt"

func TabularOutput() {

	fmt.Println("N\tN2\tN3\tN4")

	for number := 1; number <= 5; number++ {

		n := number
		n2 := number * number
		n3 := number * number * number
		n4 := number * number * number * number

		fmt.Printf("%d\t%d\t%d\t%d\n", n, n2, n3, n4)
	}
}
