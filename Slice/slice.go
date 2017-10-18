package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	var b [3]int

	b = a
	b[0] = 9

	fmt.Println("#1 Array")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	c := [...]int{1, 2, 3}
	var d [3]int
	e := &c

	for i := range d {
		d[i] = e[i]
	}

	d[0] = 9

	fmt.Println("#2 Pointer Array")
	fmt.Println(e)
	fmt.Println(d)
	fmt.Println()

	f := []int{1, 2, 3}
	var g []int

	g = f
	g[0] = 9

	fmt.Println("#3 Slice")
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println()

	f2 := []int{1, 2, 3}
	g2 := make([]int, 2, 3)

	copy(g2, f2)
	g2[0] = 9

	fmt.Println("#4 Copy Slice")
	fmt.Println(f2)
	fmt.Println(g2)
	fmt.Println()
}
