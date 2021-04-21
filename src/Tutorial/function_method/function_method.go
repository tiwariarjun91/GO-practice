package main

import("fmt")

type User struct{
	 bornYear int
	 currentYear int
}

func(u *User) Age(y int, c int) int{
	return c-y
}

func Hello(name string) (string){
	str :="Hello "+ name + " Welcome to method vs functions"
	return str

}

func main(){
	var u User
	age := u.Age(1997,2021)
	fmt.Println("Age of u is ",age)
	response:= Hello("Arjun")
	fmt.Println(response)
}