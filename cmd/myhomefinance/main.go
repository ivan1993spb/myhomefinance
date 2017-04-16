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
	parser.AddIDs = true

	repository, _ := memoryrepository.NewTransactionsRepository()
	c := core.New(repository)

	for {
		t, err := parser.ReadTransaction()
		if err != nil {
			break
		}
		c.CreateTransaction(t)
	}

	unixTime := time.Unix(0, 0)
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// All time
	inflow, outflow, balance, count := c.GetStatsByTimeRange(unixTime, now)
	fmt.Print("# Stats\n\n")
	fmt.Printf("## All time\n\n* inflow: %0.2f\n* outflow: %0.2f\n* balance: %0.2f\n* transactions: %d\n", inflow, outflow, balance, count)
	fmt.Println()

	// Month
	inflow, outflow, profit, count := c.GetStatsByTimeRange(monthStart, now)
	fmt.Printf("## %s\n\n* inflow: %0.2f\n* outflow: %0.2f\n* profit: %0.2f\n* transactions: %d\n", monthStart.Month(), inflow, outflow, profit, count)
}
