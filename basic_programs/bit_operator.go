package main

import (
	"fmt"
)

func main() {
	a := 3 
	b := 4
	c := 8
	fmt.Println(a & b)
	fmt.Println(a ^ b)
	fmt.Println(a | b)
	fmt.Println(a &^ b)
	fmt.Println(c >> 3)
	fmt.Println(c << 3)

}