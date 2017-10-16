package main

import (
	"fmt"
)

func main() {
	for i, j := 1, 0; i < 10; i, j = j+2, i-1 {
		fmt.Println(i, j)
	}
}
