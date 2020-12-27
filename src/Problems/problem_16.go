// Reading standard input
// Problem 16

package main
import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a string")
	read, _ := reader.ReadString('\n')
	fmt.Println(read, "input")
}