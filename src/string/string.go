package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Implicitly typed string"
	fmt.Printf("str1 :- %v: %T\n", str1, str1)

	// fmt.Println does not print the verbs it considers everything as a string

	var str2 string 
	str2 = "Explicitly typed string"
	fmt.Printf("str2 :- %v: %T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.Title(str1))

	lcase := "hello"
	ucase := "HELLO"
	fmt.Printf("Equal ?\n",(lcase == ucase))
	fmt.Printf("\nEqual non-case sensitive?",strings.EqualFold(lcase, ucase))
	fmt.Printf("Contains ?",strings.Contains(str1, "typ"))
	fmt.Printf("Contains ?",strings.Contains(str1, "typeds"))
	// when p is in lower case for Printf you will get the following error cannot refer to unexported name fmt.printf





}