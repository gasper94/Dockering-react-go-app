package main

import (
	"time"
	// "math/rand"
)

type Account struct {
	ID 					int			`json:"id"`
	FirstName 			string		`json:"firstName"`
	LastName 			string		`json:"lastName"`
	EncryptedPassword 	string 		`json:"-"`
	CreatedAt time.Time  			`json:createdAt`
}


func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: 			firstName,
		LastName:  			lastName,
		CreatedAt:  		time.Now().UTC(),
	}
}
