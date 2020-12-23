// extracting substrings

package main

import(
	"fmt"
	"strings"
	
)

func main(){
	mainString := "My name is Arjun"
	subString := "name"
	index := strings.Index(mainString, subString)
	extractedString := mainString[index:len(subString)+index]
	fmt.Println("Index", index)
	fmt.Println("Length of subString", len(subString))
	fmt.Println(extractedString)





	
}