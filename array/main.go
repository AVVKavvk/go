package main

import "fmt"

func main() {
	fmt.Println("Welcome to array lecture")

	var arrayList[5]string ;
	arrayList[0] = "Hello"
	arrayList[1] = "World"
	arrayList[2] = "Go"
	arrayList[4] = "Lang"

	fmt.Println("arraylist is : " , arrayList);
	fmt.Println("arraylist is : " , len(arrayList));
	
	
	newList:=[7] string{"vipin","kumawat"};
	fmt.Println("newList is : " , newList);
	fmt.Println("arraylist is : " , len(newList));
}