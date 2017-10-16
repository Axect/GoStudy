package main

import (
	"fmt"
)

func main() {
	num := 1 + 2i
	fmt.Println(num)
	fmt.Println(real(num))
	fmt.Println(imag(num))
}
