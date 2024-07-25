package main

import "fmt"

func main() {
	fmt.Println("welcome to defer code")

	defer fmt.Println("wolrd");
	defer fmt.Println("one");
	defer fmt.Println("two");
	fmt.Println("hello");
	deferDemo()
	fmt.Println("")
}

func deferDemo()  {
	
	for i := 0; i < 6; i++ {
		defer fmt.Printf("%d ",i);
	}
	
}