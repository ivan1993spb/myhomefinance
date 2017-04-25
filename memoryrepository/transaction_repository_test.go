package memoryrepository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/models"
)

func TestTransactionRepository_CreateTransaction(t *testing.T) {
	r, err := newTransactionRepository()
	require.Nil(t, err)
	require.Len(t, r.transactions, 0)
	transaction := &models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	require.Nil(t, r.CreateTransaction(transaction))
	require.Len(t, r.transactions, 1)
	require.Equal(t, transaction, r.transactions[0])
}

func TestTransactionRepository_UpdateTransaction(t *testing.T) {
	r, err := newTransactionRepository()
	require.Nil(t, err)
	require.Len(t, r.transactions, 0)
	transaction := &models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	r.transactions = []*models.Transaction{transaction}
	require.Len(t, r.transactions, 1)
	var amount float64 = 100
	require.Nil(t, r.UpdateTransaction(&models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   amount,
		Title:    "Title",
		Category: "Category",
	}))
	require.Len(t, r.transactions, 1)

	require.Equal(t, transaction.ID, r.transactions[0].ID)
	require.Equal(t, transaction.Time, r.transactions[0].Time)
	require.Equal(t, amount, r.transactions[0].Amount)
	require.Equal(t, transaction.Title, r.transactions[0].Title)
	require.Equal(t, transaction.Category, r.transactions[0].Category)
}

func TestTransactionRepository_DeleteTransaction(t *testing.T) {
	r, err := newTransactionRepository()
	require.Nil(t, err)
	require.Len(t, r.transactions, 0)
	transaction := &models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	r.transactions = []*models.Transaction{transaction}
	require.Len(t, r.transactions, 1)
	require.Nil(t, r.DeleteTransaction(transaction))
	require.Len(t, r.transactions, 0)
}

func TestTransactionRepository_GetTransactionsByTimeRange(t *testing.T) {
	var (
		beforeBefore = time.Unix(1, 0)
		before       = time.Unix(2, 0)
		bitBefore    = time.Unix(3, 4)
		current      = time.Unix(3, 5)
		bitAfter     = time.Unix(3, 6)
		after        = time.Unix(4, 0)
		afterAfter   = time.Unix(5, 0)
	)
	r, err := newTransactionRepository()
	require.Nil(t, err)
	require.Len(t, r.transactions, 0)
	transaction := &models.Transaction{
		ID:       1,
		Time:     current,
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	r.transactions = []*models.Transaction{transaction}
	require.Len(t, r.transactions, 1)

	// ------------------------------------------------------------------------

	var transactions []*models.Transaction

	transactions, err = r.GetTransactionsByTimeRange(beforeBefore, before)
	require.Nil(t, err)
	require.Len(t, transactions, 0)

	transactions, err = r.GetTransactionsByTimeRange(before, bitBefore)
	require.Nil(t, err)
	require.Len(t, transactions, 0)

	transactions, err = r.GetTransactionsByTimeRange(bitAfter, after)
	require.Nil(t, err)
	require.Len(t, transactions, 0)

	transactions, err = r.GetTransactionsByTimeRange(after, afterAfter)
	require.Nil(t, err)
	require.Len(t, transactions, 0)

	// ------------------------------------------------------------------------

	transactions, err = r.GetTransactionsByTimeRange(bitBefore, bitAfter)
	require.Nil(t, err)
	require.Len(t, transactions, 1)

	transactions, err = r.GetTransactionsByTimeRange(before, after)
	require.Nil(t, err)
	require.Len(t, transactions, 1)

	transactions, err = r.GetTransactionsByTimeRange(beforeBefore, afterAfter)
	require.Nil(t, err)
	require.Len(t, transactions, 1)
}

func TestTransactionRepository_GetTransactionsByTimeRangeCategories(t *testing.T) {
	// todo add test
}

func TestTransactionRepository_GetStatsByTimeRange(t *testing.T) {
	// todo add test
}
