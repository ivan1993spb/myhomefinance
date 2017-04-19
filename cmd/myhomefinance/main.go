package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/imports"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func main() {
	parser := imports.SimpleCSVParser{
		Reader: csv.NewReader(os.Stdin),
	}
	parser.Reader.FieldsPerRecord = 4
	parser.Reader.Comma = ';'
	parser.Reader.LazyQuotes = true
	parser.AddIDs = false

	transactionsRepository, _ := memoryrepository.NewTransactionsRepository()
	accountsRepository, _ := memoryrepository.NewAccountRepository()
	c := core.New(transactionsRepository, accountsRepository)
	account, _ := c.CreateAccount()
	if err := accountsRepository.CreateAccount(account); err != nil {
		fmt.Println(err)
		return
	}

	for {
		t, err := parser.ReadTransaction()
		if err != nil {
			break
		}
		t.AccountID = account.ID
		if err := transactionsRepository.CreateTransaction(t); err != nil {
			fmt.Println(err)
			return
		}
	}

	unixTime := time.Unix(0, 0)
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// All time
	inflow, outflow, balance, count := c.GetAccountStatsByTimeRange(account.ID, unixTime, now)
	fmt.Print("# Stats\n\n")
	fmt.Printf("## All time\n\n* inflow: %0.2f\n* outflow: %0.2f\n* balance: %0.2f\n* transactions: %d\n", inflow, outflow, balance, count)
	fmt.Println()

	// Month
	inflow, outflow, profit, count := c.GetAccountStatsByTimeRange(account.ID, monthStart, now)
	fmt.Printf("## %s\n\n* inflow: %0.2f\n* outflow: %0.2f\n* profit: %0.2f\n* transactions: %d\n", monthStart.Month(), inflow, outflow, profit, count)
}
