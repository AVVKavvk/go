package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	smg := "welcome to new web"

	fmt.Println(smg)

	reader:= bufio.NewReader(os.Stdin);
	fmt.Print("Enter your name: ")
	input,_:= reader.ReadString('\n');
	fmt.Println("Hello ", input);
}