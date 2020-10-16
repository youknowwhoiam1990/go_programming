package main

import  "fmt"

func main(){

	x:= 4
	a := &x

	fmt.Println(a)
	fmt.Println(*a)

	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)

}