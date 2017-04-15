package imports

import (
	"encoding/csv"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

const (
	simpleCSVFieldTime = iota
	simpleCSVFieldAmount
	simpleCSVFieldTitle
	simpleCSVFieldCategory
)

type SimpleCSVParser struct {
	counter uint64
	*csv.Reader
}

func NewSimpleCSVParser() *SimpleCSVParser {
	return nil
}

func (p *SimpleCSVParser) Error() error {
	return nil
}

type errReadTransaction string

func (e errReadTransaction) Error() string {
	return "cannot read transaction: " + string(e)
}

func (p *SimpleCSVParser) ReadTransaction() (*models.Transaction, error) {
	row, err := p.Read()
	if err != nil {
		return nil, errReadTransaction(err.Error())
	}

	transactionTime, err := time.Parse(time.RFC1123Z, row[simpleCSVFieldTime])
	if err != nil {
		return nil, errReadTransaction(err.Error())
	}

	amount, err := strconv.ParseFloat(row[simpleCSVFieldAmount], 64)
	if err != nil {
		return nil, errReadTransaction(err.Error())
	}

	var (
		ID       = atomic.AddUint64(&p.counter, 1)
		title    = row[simpleCSVFieldTitle]
		category = row[simpleCSVFieldCategory]
	)

	return &models.Transaction{
		ID:       ID,
		Time:     transactionTime,
		Amount:   amount,
		Title:    title,
		Category: category,
	}, nil
}
