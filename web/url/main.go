package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://vipinnotes.onrender.com:3000/notes?sem=1&subject=fcs"

func main() {
	fmt.Println("welcome to url code")
	res,err:=url.Parse(myUrl);

	if err!=nil{
		panic(err);
	}
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Port())
	// fmt.Println(res.Hostname())
	// fmt.Println(res.Path);
	// fmt.Println(res.RawQuery);

	rawQ:=res.Query();

	fmt.Printf("type is %T\n",rawQ)
	// fmt.Println(rawQ["sem"])
	// fmt.Println(rawQ["subject"])

	// for key,val:=range rawQ{
	// 	fmt.Printf("key is %v and value is %v\n",key,val)
	// }


	createUrl:=url.URL{
		Scheme: "https",
		Host: "git-scm.com",
		Path: "docs/git-checkout",
	}
	newUrl:=createUrl.String();
	fmt.Println(newUrl)
}