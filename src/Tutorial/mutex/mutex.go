package main // go cannot run non main package

import(
	"fmt"
	"sync"
)

var(
	balance int
	mut sync.Mutex
)

func Deposit(value int, wg *sync.WaitGroup ){
	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("Depositing %v amount in account with balance %v\n",value,balance)
	balance += value
	wg.Done()
}

func Withdraw(value int, wg *sync.WaitGroup ){
	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("Withdrawing %v amount in account with balance %v\n",value,balance)
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

	fmt.Printf("Final balance amount is %v \n",balance)
}