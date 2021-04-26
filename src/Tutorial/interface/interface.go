package main

import(
	"fmt"
)

type Perimeter interface { // interface should be in lower case
	 Add() float64

}

type Area interface {
	Compute() float64
} 

type Triangle struct {
	 side1,base,side3, height float64
}

type Square struct {
	 side float64
}

type Rectangle struct {
	 side1, side2 float64 //no need to write var keyword here
}

func(t Triangle) Add() float64{
	return t.side1 + t.base + t.side3
}

func(t Triangle) Compute() float64{
	return 0.5 * t.base * t.height
}

func(s Square) Add() float64 {
	return s.side*4
}

func(s Square) Compute() float64{
	return s.side * s.side
}

func(r Rectangle) Add() float64{
	return 2*r.side1 + 2*r.side2
}

func(r Rectangle) Compute() float64{
	return r.side1 * r.side2
}

func main(){
	var a,b,c Perimeter
	var h,i,j Area
	g := Square{1}
	k := Square{4}

	a = Triangle{1,2,3,4}
	h = Triangle{1,2,3,4}
	b = g
	i = k
	
	c =Rectangle{
		side1 : 1,
		side2 : 4, // , is required at the end also
	}

	j = Rectangle{
		side1 : 1,
		side2 : 4, // , is required at the end also
	}

	d:= a.Add()
	e:= b.Add()
	f:= c.Add()
	l:= h.Compute()
	m:= i.Compute()
	n:= j.Compute()
	
	fmt.Println(d,e,f)
	fmt.Println(l,m,n)


}