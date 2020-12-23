// Problem 9
// matching a regular expression with a string

package main 

import(
	"fmt"
	"regexp"
)

func main(){
	str1 := "Hello Hello World, Hello Hello Hello World !!!!"

	pattern1:= "He"
	result, _ := regexp.MatchString(pattern1, str1)
	fmt.Println(result)
}