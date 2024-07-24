package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to time lectures")
	presentTime:= time.Now();

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday "))
	createDate:=time.Date(2023,time.February,8,21,23,21,0,time.UTC);
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}