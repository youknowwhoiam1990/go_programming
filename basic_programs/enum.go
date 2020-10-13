package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main(){

	file_Size := 4000000000.
	fmt.Printf("%.2fGB", file_Size/GB)// output: 3.73GB




}