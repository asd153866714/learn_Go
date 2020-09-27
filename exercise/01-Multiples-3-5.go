package main

import (
	"fmt"
)

func findMultiples() int {
	var multipleSum int

	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multipleSum += i
		}

	}

	return multipleSum
}

func main() {
	fmt.Println(findMultiples())
}
