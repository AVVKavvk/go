package main

import "fmt"

func main() {

	fmt.Println("Welcome to remove element code");

	cources:=[]string{"DSA","CP","OOPS","CN","OS","DBMS"};

	fmt.Println(cources);
	index:=2;
	cources=append(cources[:index],cources[index+1:]... );
	fmt.Println(cources);

}