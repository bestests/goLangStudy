package main

import (
	"fmt"
)

func superAdd (numbers ...int) (result int) {

	result = 0

	for idx, number := range numbers {

		result += number

		fmt.Println("idx - ", idx, ", number - ", number, ", result - ", result)
	}

	return
}

func main () {
	superAddResult := superAdd(1, 2, 3, 4, 5)

	fmt.Println("superAddResult - ", superAddResult)
}