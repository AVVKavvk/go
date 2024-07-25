package main

import "fmt"

func main() {
	greeter();

	result:=adder(3,5);

	fmt.Println(result)

	proResult,message:=proAdder(1,2,3,4,5,6,7,8,9,10);
	fmt.Println(proResult)
	fmt.Println(message)
}

func greeter() {
	fmt.Println("Namaste from vipin")
}

func adder(val1 int,val2 int) int{

	return val1+val2;
}

func proAdder(values ...int) (int,string)  {
	total:=0;
	for _,val:=range values{
		total+=val;
	}

	return total,"HI this is vipin's code";
}