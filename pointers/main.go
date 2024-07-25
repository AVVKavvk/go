package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers lectures")

	number := 23
	// fmt.Println(number)
	fmt.Println(&number)

	// ptr:=&number
	// fmt.Println(ptr)

	var ptr *int=&number;
	fmt.Println(ptr)

	fmt.Println(*ptr)
	*ptr=*ptr*3;
	fmt.Println(*ptr)
	fmt.Println(number)
}