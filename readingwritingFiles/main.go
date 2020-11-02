package main

import (

	"fmt"
	"io/ioutil"
 )






func main() {



	data, err := ioutil.ReadFile("sample.data")

	if err != nil {

		fmt.Println(err)
	}


	fmt.Println(string(data))




	newdata := []byte("this is the new data")


	err2 := ioutil.WriteFile("newdata.data", newdata, 0777)


	if err2 != nil{
		fmt.Println(err2)
	}

	NewData,err2 := ioutil.ReadFile("newdata.data")

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(NewData))



}