package main

import (

	"fmt"
	"bufio"
	"os"
	//"strconv"
	//"strings"

)

func main(){

	var s string
	fmt.Println("Please enter the value of s")
	fmt.Scanln(&s)
	fmt.Println(s)


	//using reader object

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some text")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)





}