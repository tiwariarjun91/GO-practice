// converting a structure to a json format
// Problem 25

package main

import(
	"fmt"
	"encoding/json"
)

type Data struct{
	Name, Age string // need to export by making the first letter to uppercase so that the json package can see it
}

func main(){

	var user Data
	user.Name = "Arjun Tiwari"
	user.Age = "23"
	fmt.Println(user)
	user1 := Data{Name : "Parth Tiwari"}
	user1.Age = "24"

	jsonContent1, err := json.Marshal(user)
	if err != nil{
		panic(err)
	}

	jsonContent2, err := json.Marshal(user1)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(jsonContent1))
	fmt.Println(string(jsonContent2))

}