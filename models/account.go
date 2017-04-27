package models

import (
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/iso4217"
)

type Account struct {
	UUID     uuid.UUID
	UserUUID uuid.UUID
	Name     string
	Currency iso4217.Currency
}
