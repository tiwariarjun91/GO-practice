// Problem 27
// getting a json content to a string content

package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	url := "http://services.explorecalifornia.org/json/tours.php"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	jsonContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	stringContent := string(jsonContent)
	fmt.Println(stringContent)
}