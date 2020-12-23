// Removing elements from the map
package main

import(
	"fmt"
)

func main(){

	indianPresident := make(map[string]string)

	indianPresident["1"] = "Rajendra Prasad"
	indianPresident["2"] = "Sarvepalli Radhakrishnan"
	indianPresident["3"] = "Zakir Husain"
	indianPresident["4"] = "Varaha Giri Venkata Giri"
	indianPresident["5"] = "Fakhrauddin Ali Ahmed"
	indianPresident["6"] = "Neelam Sanjeeva Reddy"
	indianPresident["7"] = "Zail Singh"
	indianPresident["8"] = "Ramaswamy Venkatraman"
	indianPresident["9"] = "Shankar Dayal Sharma"
	indianPresident["10"] = "Kocheril Raman Narayanan"
	indianPresident["11"] = "Dr. APJ Abdul Kalam Azad"
	indianPresident["12"] = "Dr. Pratibha Patil"
	indianPresident["13"] = "Pranab Mukherjee"
	indianPresident["14"] = "Ram Nath Kovid"

	fmt.Println(indianPresident)

	var decision string 
	fmt.Println("Do you wish to remove any president from the list?")
	fmt.Scanln(&decision)

	for decision=="yes" {
		var key string
		fmt.Println("Please enter the number of president you want to remove")
		fmt.Scanln(&key)

		delete(indianPresident, key)
		fmt.Println("Updated list")
		for i := range indianPresident{
			fmt.Println("President no ", i," :- ",indianPresident[i])
		}
		fmt.Println("Do you wish to remove any president from the list?")
		fmt.Scanln(&decision)
	}


}
