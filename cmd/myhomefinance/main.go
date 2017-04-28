package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/imports"
	"github.com/ivan1993spb/myhomefinance/iso4217"
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

	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	user, _ := c.CreateUser()
	account, _ := c.CreateAccount(user.UUID, "main", iso4217.RUB)

	for {
		t, err := parser.ReadTransaction()
		if err != nil {
			break
		}
		t.UserUUID = user.UUID
		t.AccountUUID = account.UUID
		if err := transactionRepository.CreateTransaction(t); err != nil {
			fmt.Println(err)
			return
		}
	}

	unixTime := time.Unix(0, 0)
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	fmt.Print("# Stats\n\n")

	// All time
	stats, _ := c.GetUserAccountStatsByTimeRange(user.UUID, account.UUID, unixTime, now)
	fmt.Printf("## All time\n\n* inflow: %0.2f\n* outflow: %0.2f\n* balance: %0.2f\n* transactions: %d\n",
		stats.Inflow, stats.Outflow, stats.Profit, stats.Count)
	fmt.Println()

	// Month
	stats, _ = c.GetUserAccountStatsByTimeRange(user.UUID, account.UUID, monthStart, now)
	fmt.Printf("## %s\n\n* inflow: %0.2f\n* outflow: %0.2f\n* profit: %0.2f\n* transactions: %d\n", monthStart.Month(),
		stats.Inflow, stats.Outflow, stats.Profit, stats.Count)
	fmt.Println()

	fmt.Print("## Month cotegories\n\n")
	categoriesSums, _ := c.CountUserAccountCategorySumsByTimeRange(user.UUID, account.UUID, monthStart, now)

	if len(categoriesSums) > 0 {
		fmt.Printf("| %30s | %6s | %15s |\n", "Category", "Count", "Sum")
		fmt.Println("|:-------------------------------|:-------|:----------------|")
		for _, categorySum := range categoriesSums {
			fmt.Printf("| %30s | %6d | %15.2f |\n", categorySum.Category, categorySum.Count, categorySum.Sum)
		}
	} else {
		fmt.Println("empty")
	}
	fmt.Println()
}
