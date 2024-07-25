package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["py"] = "python"
	languages["go"] = "golang"
	languages["rb"] = "ruby"

	// fmt.Println("all languages are :")
	// fmt.Println(languages);

	// for key,val:=range languages{
	// 	fmt.Printf("key is %s and value is %s\n",key,val);
	// }

	fmt.Println(languages["js"])
	fmt.Println(languages["ja"])
}