package main

import(
	"fmt"
)

type Person struct{
	Name string
	Age int



}


func main(){
	arjun := Person{"Arjun Nandkishor Tiwari", 22}
	parth := Person{"Parth Jitendra Tiwari", 24}

	fmt.Println(arjun, "\n", parth)
	fmt.Printf("Name :- %v\nAge :- %v", arjun.Name, arjun.Age)
}