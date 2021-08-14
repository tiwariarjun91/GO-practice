package main

import(
	"fmt"
	"sync"
)

var (
	balance int
	mutex sync.Mutex
)

func Deposit(value int, wg *sync.WaitGroup){
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("value: %v, balance: %v, Operaiton: Deposit",value,balance)
	balance += value
	wg.Done()
}

func Withdraw(value int, wg *sync.WaitGroup){
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("value: %v, balance: %v, Operaiton: Withdraw",value,balance)
	balance -= value
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	balance = 1000
	wg.Add(2)
	Withdraw(700,&wg)
	Deposit(500,&wg)
	wg.Wait()
	fmt.Printf("Final balance: %v",balance)

}