package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := int64(100)
	numberStr := strconv.FormatInt(number, 10)
	fmt.Println(numberStr)
	number2 := 100
	numberStr2 := strconv.Itoa(number2)
	fmt.Println(numberStr2)

	numberFloat := 234445221.122333
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', 2, 64)
	//numberFloatStr := strconv.FormatFloat(numberFloat, 'f', -1, 64)
	fmt.Println(numberFloatStr)

}
