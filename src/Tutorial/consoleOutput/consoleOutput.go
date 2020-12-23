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

	// if you do not want to address any variable use _ in place of it's name

	stringLength2, _ := fmt.Println(str1, str2, str3)
	fmt.Println("String length 2:-", stringLength2)

	// using verbs i.e %v in fmt.Printf func

	aNumber := 23
	isTrue := true

	fmt.Printf("Value of aNumber is:- %v\n",aNumber)
	fmt.Printf("Value of isTrue is:- %v\n",isTrue)
	fmt.Printf("Value of aNumber upto 2 decimal is:- %.2f\n", float64(aNumber))
	fmt.Printf("Data types are, %T, %T, %T, %T, %T, %T", str1, str2, str3, aNumber, isTrue, float64(aNumber))
	stringvar := fmt.Sprintf("\n specific Data types are, %T, %T, %T, %T, %T, %T", str1, str2, str3, aNumber, isTrue, float64(aNumber))
	fmt.Println(stringvar)

}

// learned about the fmt.Println function and it's return value
// different verbs 
// for float 0.f and float64(variablename)
// fmt.Sprintf("")// assigns to a string variable