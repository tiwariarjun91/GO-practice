// converting time to string and string to time

package main 

import(
	"fmt"
	"time"
)

func main(){
	s := "1997-Dec-24"
	const layout = "2020-Jan-23"
	t, _ := time.Parse(layout,"1997-Dec-24")

	fmt.Println("original string :- ", s)
	fmt.Println("Formated time :- ", t)
}