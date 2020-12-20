package main

import (
	"fmt"
	"math/big"
	"math"
)

func main(){

	var b1, b2, b3, bigsum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(53.5)
	b3.SetFloat64(121.45)
	bigsum.Add(&b1,&b2).Add(&b3, &bigsum)
	fmt.Println(&b1, &b2, &b3, "the sum is ", &bigsum)
	b4 := 45.5
	fmt.Println(b4)
	var b5 big.Float
	b5.SetFloat64(45.2)
	fmt.Println(&b5)

	// need to check whether big.Float creates pointers as printing b1 instead of &b1 results in vague output and need to check the use of SetFloat64

	circleRadius := 45.2
	circumference := 2*circleRadius*(math.Pi)
	fmt.Printf("circumference is %0.2f", circumference)

	// need to read about when to use & and when not to


}