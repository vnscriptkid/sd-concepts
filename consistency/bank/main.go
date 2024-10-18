package main

import (
	"fmt"
	"sync"
)

// Account represents a bank account
type Account struct {
	mutex   sync.Mutex
	balance int
}

// Deposit adds money to the account
func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	defer a.mutex.Unlock()

	fmt.Printf("Depositing %d to account\n", amount)
	a.balance += amount
}

// Withdraw removes money from the account
func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.balance >= amount {
		fmt.Printf("Withdrawing %d from account\n", amount)
		a.balance -= amount
	} else {
		fmt.Println("Insufficient funds")
	}
}

// Balance returns the current balance
func (a *Account) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.balance
}

func main() {
	account := &Account{balance: 1000}
	var wg sync.WaitGroup

	wg.Add(2)
	go account.Deposit(500, &wg)
	go account.Withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("Final Balance: %d\n", account.Balance())
}
