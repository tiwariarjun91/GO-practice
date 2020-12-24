// Problem 13
// merging two maps

package main

import(
	"fmt"
)

func main(){
	map1 := make(map[int]string)
	map2 := make(map[int]string)

	map1[1] = "A"
	map1[2] = "B"
	map2[3] = "C"
	map2[4] = "D"

	for i := range map2{
		map1[i] = map2[i]
	} 
	fmt.Println("Map 1:- ", map1)
	fmt.Println("Map 2:- ", map2)

}