// Problem 29
// adding one month to a date and getting exact date

package main

import(
	"fmt"
	"time"
)

func main(){
	date := time.Date(2020, time.January, 30, 0, 0, 0, 0, time.UTC)

	days_30 := date.AddDate(0, 0, 30)

	fmt.Println(date)
	//fmt.Println("\n")
	fmt.Println("date + 30 days")
	//fmt.Println("\n")
	fmt.Println(days_30)

}