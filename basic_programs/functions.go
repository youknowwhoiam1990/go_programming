package main

import "fmt"


func add(num1,num2 float64) (float64){ //types of variables passed and ttype pf return value must be given explicitly
	return num1 + num2
}

func multi(num1, num2 float64) float64{
	return num1*num2
}
func str(a,b string) (string){
	return a+ " " + b
} 
func main(){
	num1, num2 := 3.5,7.8
	a,b :=  "Hello","World"
	fmt.Println(add(num1,num2))
	fmt.Println(multi(num1,num2))
	fmt.Println(str(a,b))
}