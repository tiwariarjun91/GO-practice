package main

import (

	"fmt"

)

func main(){

	str1:= "Hello Arjun "
	str2:= "How you doin"
	str3:= " I hope great"

	fmt.Println(str1, str2, str3)

	stringLength,error := fmt.Println("this is to check the count of the string")

	if error == nil {
		fmt.Println(stringLength)
	}


}

// learned about the fmt.Println function and it's return value