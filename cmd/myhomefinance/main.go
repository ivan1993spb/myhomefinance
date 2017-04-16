package main

import (
	"encoding/csv"
	"fmt"
	"os"

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
		t, _ := parser.ReadTransaction()
		if t == nil {
			break
		}
		c.CreateTransaction(t)
	}

	inflow, outflow, balance := c.GetStats()
	fmt.Printf("All time\ninflow: %0.2f\noutflow: %0.2f\nbalance: %0.2f\n", inflow, outflow, balance)
	fmt.Println("----")
	inflow, outflow, profit := c.GetStatsMonth()
	fmt.Printf("Month\ninflow: %0.2f\noutflow: %0.2f\nprofit: %0.2f\n", inflow, outflow, profit)
}
