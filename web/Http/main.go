package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://vipinnotes.onrender.com/"

func main() {

	res, err := http.Get(url);

	if err!=nil{
		panic(err)
	}
	fmt.Printf("type of  %T\n",res);
	fmt.Printf("status code %d\n",res.StatusCode);

	defer res.Body.Close();

	dataByte,err:=ioutil.ReadAll(res.Body);
	if err!=nil{
		panic(err)
	}
	// fmt.Printf("type of dataByte %T\n",dataByte);
	// fmt.Printf("dataByte %s\n",dataByte);
	fmt.Println(string(dataByte))
}