package main 

import (

	"fmt"
)

const a int16 = 12 // this constant is accessible by the main function
func main(){

	const a = 42
	var b int16 = 27
	fmt.Printf("%v,%T\n",a,a) //42, int inner constnat shadows the package level constant.
	fmt.Printf("%v,%T\n",a+b, a+b) // 69, int16, implicit converison of int a to int16


}