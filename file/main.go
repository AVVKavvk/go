package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to file read and write code")
	
	file,err:=os.Create("vipin.txt");
	content:="hey vipin can code in go "
	checkNilErr(err)
	length,err:=io.WriteString(file,content);
	checkNilErr(err)
	fmt.Printf("length of the file is %d\n",length);
	defer file.Close();

	readFile("./vipin.txt");
}

func readFile(fileName string){
	data,err:=ioutil.ReadFile(fileName);
	checkNilErr(err)
	fmt.Println(string(data));
}


func checkNilErr(err error){
	if err!=nil{
		panic(err);
	}
}