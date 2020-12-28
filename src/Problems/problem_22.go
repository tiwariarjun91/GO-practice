// Problem 22
// using container package to implement list, heaps etc

package main

import(
	"fmt"
	"container/list"
	//"container/heap"
)

func main(){
	l := list.New()
	e1 := l.PushFront(1)
	e4 := l.PushBack(4)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	fmt.Println(l)
	for i := l.Front(); i!= nil; i = i.Next(){
		fmt.Println(i.Value)
	}
}