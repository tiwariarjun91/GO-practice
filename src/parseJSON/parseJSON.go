package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	url := "http://services.explorecalifornia.org/json/tours.php"

	content := contentFromServer(url)

	fmt.Print(content)

}

func contentFromServer(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	} 
	contentJSON, err := ioutil.ReadAll(resp.Body) // if we do not pass resp.Body then it gives error
	if err != nil {
		panic(err)
	}

	contentString := string(contentJSON)

	return contentString
}