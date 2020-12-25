// creating a map of custom type
// Problem 24

package main

import(
	"fmt"
)

type Data struct{
	 FirstName string
	 LastName string
}

func main(){
	n1 := Data{FirstName: "Arjun", LastName: "Tiwari"}
	n2 := Data{FirstName: "Sarang", LastName: "Girdhari"}
	n3 := Data{FirstName: "Omar", LastName: "Khan"}

	data := make(map[int]Data)
	data[1] = n1
	data[2] = n2
	data[3] = n3

	fmt.Println("map with a custom struct value type :- ", data)
}