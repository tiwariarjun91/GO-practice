package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Tour struct{
	Name, Price string
}

func main(){
	url := "http://services.explorecalifornia.org/json/tours.php"

	content := contentFromServer(url)

	tours := toursfromJSON(content)

	fmt.Print(tours)

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

func toursfromJSON(content string) []Tour{
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()

	checkError(err)

	var tour Tour
	for decoder.More(){
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours



}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}