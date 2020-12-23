//Problem 1
package main

import(
	"fmt" // you sometimes refer it as lmt 
)

func main(){


	indianPresident := make(map[int]string)

	indianPresident[1] = "Rajendra Prasad"
	indianPresident[2] = "Sarvepalli Radhakrishnan"
	indianPresident[3] = "Zakir Husain"
	indianPresident[4] = "Varaha Giri Venkata Giri"
	indianPresident[5] = "Fakhrauddin Ali Ahmed"
	indianPresident[6] = "Neelam Sanjeeva Reddy"
	indianPresident[7] = "Zail Singh"
	indianPresident[8] = "Ramaswamy Venkatraman"
	indianPresident[9] = "Shankar Dayal Sharma"
	indianPresident[10] = "Kocheril Raman Narayanan"
	indianPresident[11] = "Dr. APJ Abdul Kalam Azad"
	indianPresident[12] = "Dr. Pratibha Patil"
	indianPresident[13] = "Pranab Mukherjee"
	indianPresident[14] = "Ram Nath Kovid"

	//fmt.Println(indianPresident)

	// Normal iteration
	fmt.Println("\n Normal Iteration\n")

	for i:=1;i<11;i++ { //using range function resulted in random genereation of all the values 
		fmt.Println("President no ",i," :-", indianPresident[i],"\n")
	}

	// Iteration witht the help of range function
	fmt.Println("\nIteration with a range function\n")
	for j:= range indianPresident{
		fmt.Println("President no ",j," :-", indianPresident[j],"\n")
	}

	// Iteration with both the key and value
	fmt.Println("\nIteration with a range function // both key and value\n")

	for k, v:= range indianPresident{
		fmt.Println("President no ",k,":- ",v)
	}



}