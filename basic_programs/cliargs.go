package main 

import (
	"fmt"
	"flag"
)

func main(){

	var name string

	flag.StringVar(&name,"name","","Usage")

	flag.Parse()

	fmt.Println("Hello " + name)
}