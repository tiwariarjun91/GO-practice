// problem 23
// creating custom error type

package main 

import(
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("./filename.txt")
	if err == nil {
		fmt.Println("File:- ", file)
	} else{
		fmt.Println("Error :- ", err.Error())
	}
}