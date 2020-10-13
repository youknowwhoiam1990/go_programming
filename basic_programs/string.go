package main 

import (
	"fmt"
)

func main(){

	var s string = "This is go program"
	t := "This is a string"

	b := []byte(s)
	fmt.Printf("%v,%T\n",s,s) //output:This is go program,string
	fmt.Printf("%v, %T\n", t,t)//output:This is a string, string
	fmt.Printf("%v, %T\n", b,b) //output:[84 104 105 115 32 105 115 32 103 111 32 112 114 111 103 114 97 109], []uint8
}