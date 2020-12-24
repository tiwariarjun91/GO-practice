// problem 14
//merging 2 arrays

package main

import(
	"fmt"
)

func main(){
	arr1 := []int{1,2,3,4,5}
	arr2 := [5]int{6,7,8,9,10}

	for i := 0 ;i<=4 ;i++{
		arr1 = append(arr1, arr2[i])
	}
	fmt.Println("arr1 :- ", arr1)
	fmt.Println("arr2 :- ", arr2)
}