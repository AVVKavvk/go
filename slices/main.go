package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to the silces lectures")

	// firstSlices:=[]string{"vipin","kumawat"};

	// fmt.Printf("type of : %T\n",firstSlices);
	// fmt.Println(firstSlices)

	// firstSlices=append(firstSlices, "sikar" , "Rajasthan");
	// fmt.Println(firstSlices)

	// firstSlices=append(firstSlices[:]);
	// firstSlices=append(firstSlices[:3]);
	// firstSlices=append(firstSlices[1:3]);
	// fmt.Println(firstSlices)

	secondSlices:= make([]string ,4);
	secondSlices[0]="vipin";
    secondSlices[1]="kumawat";
	secondSlices[3]="vipin";
    secondSlices[2]="kumawat";
	secondSlices=append(secondSlices, "srijan","synder");
	fmt.Println(secondSlices);

}