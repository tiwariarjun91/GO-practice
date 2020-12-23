// String to Float
// Float to String
// problem 12

package main

import(
	"fmt"
	"strconv"
)

func main(){
	s1 := "45"
	f2 := 45.5
	f1, _:= strconv.ParseFloat(s1, 64)
	fmt.Println("String to Float")
	fmt.Printf("%v Type %T", f1, f1)
	s2:= strconv.FormatFloat(f2, 'E', -1, 64)
	fmt.Println("Float to String")
	fmt.Printf("%v Type %T", s2, s2)

}