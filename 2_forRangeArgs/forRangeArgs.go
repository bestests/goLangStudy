package forRangeArgs

import (
	"fmt"
)

func SuperAdd(numbers ...int) (result int) {

	result = 0

	for idx, number := range numbers {

		result += number

		fmt.Println("idx - ", idx, ", number - ", number, ", result - ", result)
	}

	return
}
