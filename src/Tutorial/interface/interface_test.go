package main

import(
	"testing"
	"fmt"
)

func TestAdd(t *testing.T){
	var a,b,c Perimeter
	a = Triangle{4,5,6,7}
	b = Square{5}
	c = Rectangle{
		side1 : 5,
		side2 : 6,
	}
	if(a.Add() != float64(15)){
		t.Fatalf("Incorrect Perimeter of Triangle")

	} else{ fmt.Println(a.Add())}
	if(b.Add() != float64(20)){
		t.Fatalf("Incorrect Perimeter of Square")
	} else{ fmt.Println(b.Add())}
	if(c.Add() != float64(5 * 2 + 6 * 2)){
		t.Fatalf("Incorrect Perimeter of Rectangle")
	} else{ fmt.Println(c.Add())}
	 
}

func TestCompute(t *testing.T){
	var a,b,c Area
	a = Triangle{4,5,6,7}
	b = Square{5}
	c = Rectangle{
		side1 : 5,
		side2 : 6,
	}
	if(a.Compute() != float64(0.5*5*7)){
		t.Fatalf("Incorrect Area of Triangle")
	} else{ fmt.Println(a.Compute())}
	if(b.Compute() != float64(5*5)){
		t.Fatalf("Incorrect Area of Square")
	} else{ fmt.Println(b.Compute())}
	if(c.Compute() != float64(5*6)){
		t.Fatalf("Incorrect Area of Rectangle")
	} else{ fmt.Println(c.Compute())}

}

