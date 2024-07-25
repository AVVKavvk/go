package main

import "fmt"

func main() {

	// days := []string{"sun", "tue", "wed", "fri", "sat"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:=range days{
	// 	fmt.Println(i," -> ",days[i])
	// }

	// for index,key:= range days{
	// 	fmt.Println(index," -> ",key)
	// }
	// for _,key:= range days{
	// 	fmt.Println(key)
	// }

	index:=1;

	for true{

		// if(index==12){
		// 	goto lco;
			
		// }
		if(index==15){
			goto lco;
		}
		if(index==12){
			break;	
		}

		if(index==9){
			index++;
			continue;
		}
		fmt.Println(index)
		index++;
	}
	lco:{
		fmt.Println("lco")
	}

	fmt.Println(index)
}