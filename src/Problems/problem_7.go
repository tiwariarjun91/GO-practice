// Problem 7
// Difference between dates

package main 

import(
	"fmt"
	"time"
	"math"
)

func main(){
	date1 := time.Date(2020, time.May, 12, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2045, time.June, 30, 0, 0, 0, 0, time.UTC)

	date_diff := math.Abs(date1.Sub(date2).Hours()/24)

	fmt.Println(date_diff)
}