package main

import "math/rand"

type Account struct {
	ID int `json:"id"`
	AccountNumber int64 `json:"account_number"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Balance float64 `json:"balance"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		AccountNumber: int64(rand.Intn(10000)),
	}
}
