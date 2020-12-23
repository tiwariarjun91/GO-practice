// replacing a substring
// Problem 10

package main 

import(
	"fmt"
	"strings"
)

func main(){
	str1:= "Hello, Hello, Hello World"
	str2:= strings.ReplaceAll(str1, "Hello", "Hola")
	fmt.Println("Original String :- ",str1 )
	fmt.Println("String after replacement :- ", str2)
}