package main

import "fmt"

func main() {

	fmt.Println("welcome to struct code")

	vipin:=User{"vipin",21};

	fmt.Println(vipin)

	fmt.Printf("%+v\n",vipin);

	fmt.Printf("name is %s and age is %v",vipin.Name,vipin.Age)
}

type User struct {
	Name string
	Age  int
}