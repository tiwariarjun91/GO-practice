// different ways to initialize arrays

package main 

import(
	"fmt"
)

func main(){
	var data  = []string{"Arjun", "Tiwari", "Pareek"}
	data1 := []string{"This is also a string array"}
	var data2 []string
	//data2[0] = "zero position" gives error
	data2 = append(data2, "Appended posiiton")
	data3 := make([]string, 5)
	data3[0] = "zeroth position"

	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)
}