package main

import ("fmt")

func main(){

	var x int = 2
	var msg string

	switch x {

	case 1:
		msg = "value is 1"
	case 2:
		msg = "value is 2"
	case 3:
		msg = "value is 3"
	case 4:
		msg = "value is 4"
	default:
		msg = "Default value"
	}

	fmt.Println(msg)



}