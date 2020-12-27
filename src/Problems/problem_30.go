// Problem 30
// Converting a string array into a byte array

package main

import(
	"fmt"
	"bytes"
	"encoding/gob"
)

func main(){
	str1 := []string{"foo", "git", "push"}
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(str1)
	bf := buf.Bytes()
	fmt.Println(bf)
	
	



}

