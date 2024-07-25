package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// rand.Seed(time.Now().UnixNano());

	value:=rand.Intn(6)+1;

	switch value{
		case 1:{
			fmt.Println("you can open your card");
			}
		case 2:{
			fmt.Println("move 2 steps");
		}	
		case 3:{
			fmt.Println("move 3 steps");
		}
		case 4:{
			fmt.Println("move 4 steps");
		}
		case 5:{
			fmt.Println("move 5 steps");
		}
		case 6:{
			fmt.Println("move 6 steps and roll dice again");
		}
		default :{
			fmt.Println("error");
		}
	}
}