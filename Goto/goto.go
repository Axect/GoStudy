package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("> ")
	fmt.Scanln(&a)

	if a == 1 {
		goto Error1
	}

	if a == 2 {
		goto Error2
	}

	if a == 3 {
		goto Error3
	}

	fmt.Println(a)
	fmt.Println("Success")
	return

Error1:
	fmt.Println("Error1")
	return

Error2:
	fmt.Println("Error2")
	return

Error3:
	fmt.Println("Error3")
	return
}
