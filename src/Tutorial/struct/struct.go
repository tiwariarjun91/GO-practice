package main

import(
	"fmt"
)

type Person struct{
	Name string // do not require var keyword while defining fields in the struct
	Age int
	Demo struct{
		Height int
		
	}



}

/*type Demo struct{
	Height int
}*/


func main(){
	arjun := Person{
		Name :"Arjun Nandkishor Tiwari", 
		Age : 22, 
		Demo:{
			Height : 175},}
	/*parth := Person{"Parth Jitendra Tiwari", 24, Demo:{Height : 174}}*/
	var p Person
	/*var q Person*/
	p.Name = "Gauresh"
	p.Demo.Height = 175

	fmt.Println(p)
	
	
	
	
	/*a := &Demo{}
	a.Height = 45
	fmt.Println(*a, "\n",)*/

	fmt.Println(arjun, "\n")
}