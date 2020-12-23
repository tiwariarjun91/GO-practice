//Problem 11
// Removing duplicate elements form an array

package main 

import(
	"fmt"
)

func main(){
	arr1 := []int{1,2,1,3,1,4,5,20,56,56,4,23,4,12,100}
	var arr2 []int

	for i:= 0;i<len(arr1);i++{
		find := Find(arr2, arr1[i])
		if find==1{
			arr2 = append(arr2, arr1[i])
		}
	}
	fmt.Println("Original array :- ", arr1)
	fmt.Println("Unique Array", arr2)
}

func Find(arr []int, a int) int {
	count:=0
	if len(arr)==0{
		return 1
	} else{
	for i:= 0; i< len(arr);i++{
		if arr[i]==a{

			count= 1
		}}
		if count==1{
			return 0
		}else{
			return 1
		}
		
	
	}
	} 


