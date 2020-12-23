package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main(){

	content:= "Sample text to check if the file reads successfully"

	file , err := os.Create("./readFile.txt")
	errorChecker(err)
	defer file.Close()
	ln, err := io.WriteString(file, content)
	errorChecker(err)
	fmt.Println(ln)


	cont, err := ioutil.ReadFile("./readFile.txt")
	errorChecker(err)

	fmt.Println("Content in bytes :-", cont)

	stringContent := string(cont)

	fmt.Println("\nContent in string :- ", stringContent)

	/*file1, err := os.Open("./readFile.txt")

	errorChecker(err)
	defer file1.Close()

	ln1, err := io.WriteString(file1, content)

	errorChecker(err)
	fmt.Println(ln1)*/



}

func errorChecker(err error){

	if err != nil{

		panic(err)
	}
}