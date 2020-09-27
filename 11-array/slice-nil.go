package main

import "fmt"

func main() {
	var s []int // 等於 s := []int{}
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// The zero value of a slice is "nil"
