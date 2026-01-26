package main

import "fmt"

// TYPE
type BankAccount struct {
    Owner  string
    Balance int
}

// METHOD
func (b *BankAccount) Deposit(amount int) {
    b.Balance += amount
}

func main() {
    acc := BankAccount{
        Owner:   "Soumya",
        Balance: 1000,
    }

    acc.Deposit(500)

    fmt.Println("Owner:", acc.Owner)
    fmt.Println("Balance:", acc.Balance)
}
