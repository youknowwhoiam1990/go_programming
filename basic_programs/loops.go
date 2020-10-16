package main

import (
	"fmt"
)

func main(){

	// for i:= 0; i<100; i++{
	// 	if i%2 == 0{
	// 		fmt.Println(i)
	// 	}
	// }

	for i:= 1; i<100; i++{
		if i%3 == 0 && i%2 == 0{
			fmt.Println(i)
		}
	}
}