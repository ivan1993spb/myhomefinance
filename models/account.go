package models

import "github.com/ivan1993spb/myhomefinance/iso4217"

type Account struct {
	ID       uint64
	Currency iso4217.Currency
}
