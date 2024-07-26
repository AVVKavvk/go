package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to goWithNode code");

	// getRequest();
	postrequest();
}

func getRequest()  {
	const url="http://localhost:3004";

	res,err:=http.Get(url);

	defer res.Body.Close();

	if err!=nil{
		panic(err)
	}

	fmt.Println("status code is :",res.StatusCode)
	fmt.Println("Content length is :",res.ContentLength)

	dataByte,err:=ioutil.ReadAll(res.Body);

	// content:=string(dataByte)

	// fmt.Println(content)
	var responseString strings.Builder;

	responseString.Write(dataByte);
	fmt.Println(responseString.String())
	fmt.Println(responseString.Len())
}

func postrequest(){
	const url="http://localhost:3004/post";

	responseBody:= strings.NewReader(`
		{
			"name":"vipin",
			"age":"21"

		}
	`)

	res,err:=http.Post(url,"application/json",responseBody);
	defer res.Body.Close();

	if err!=nil{
		panic(err)
	}
	dataByte,err:=ioutil.ReadAll(res.Body);

	var responseString strings.Builder;
	responseString.Write(dataByte);
	fmt.Println(res.StatusCode);
	fmt.Println(res.ContentLength)
	fmt.Println(responseString.String())

}