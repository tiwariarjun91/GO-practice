// type nul > webRequest.go

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	if err!= nil{
		//fmt.Println("Error : ", err)
		panic(err)
	} else{
		fmt.Printf("Response type is :- %T ",resp)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err!= nil{
		//fmt.Println("Error : ", err)
		panic(err)
	}

	contentString := string(bytes)

	fmt.Printf(contentString)

}