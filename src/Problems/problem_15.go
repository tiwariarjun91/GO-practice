// creating a temprary file
// problem 15

package main

import(
	"fmt"
	"io"
	"os"
)

func main(){

	content := "A string to be written in the file"

	file, err := os.Create("./writeString.txt")
	errorChecker(err)
	defer file.Close()
	ln, err := io.WriteString(file, content)
	errorChecker(err)
	fmt.Printf("the length is %v\n", ln)
	fmt.Println("The Length is", ln)

}



	func errorChecker(err error){

		if err != nil {
			panic(err)
		}
	}
