package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcomeMsg := "welcome to this webiste"

	fmt.Println(welcomeMsg)
    reader:=bufio.NewReader(os.Stdin);

	fmt.Println("enter value to be added one");

	input,err :=reader.ReadString('\n');

	if err!=nil{
		fmt.Println(err)
	}else{
		userValue,_:=strconv.ParseFloat(strings.TrimSpace(input),64);
		fmt.Println("valued updated to ",userValue+1)
	}
}