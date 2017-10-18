package main

import (
	"fmt"
)

func main() {
	i := 2

	switch i {
	case 1, 3, 5:
		fmt.Println("Odd")
		fallthrough
	case 2, 4, 6:
		fmt.Println("Even")
		fallthrough
	default:
		fmt.Println("Default")
	}
}
