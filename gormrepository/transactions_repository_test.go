package gormrepository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/models"
)

func TestTransactionsRepository_CreateTransaction(t *testing.T) {
	db, err := openDB()
	require.Nil(t, err)
	defer db.Close()

	r, err := newTransactionsRepository(db)
	require.Nil(t, err)
	tr := &models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	require.Nil(t, r.CreateTransaction(tr))

	dbTransaction := &transaction{}
	err = r.db.First(&dbTransaction).Error
	require.Nil(t, err)

	require.Equal(t, dbTransaction.ID, tr.ID)
	// todo check time
	require.Equal(t, dbTransaction.Amount, tr.Amount)
	require.Equal(t, dbTransaction.Title, tr.Title)
	require.Equal(t, dbTransaction.Category, tr.Category)

	err = r.db.Delete(dbTransaction).Error
	require.Nil(t, err)
}

func TestTransactionsRepository_UpdateTransaction(t *testing.T) {
	db, err := openDB()
	require.Nil(t, err)
	r, err := newTransactionsRepository(db)
	require.Nil(t, err)
	tr := &transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   0,
		Title:    "Title",
		Category: "Category",
	}
	err = r.db.Create(tr).Error
	require.Nil(t, err)

	var amount float64 = 100
	require.Nil(t, r.UpdateTransaction(&models.Transaction{
		ID:       1,
		Time:     time.Now(),
		Amount:   amount,
		Title:    "Title",
		Category: "Category",
	}))

	dbTransaction := &transaction{}
	err = r.db.First(&dbTransaction).Error
	require.Nil(t, err)

	require.Equal(t, tr.ID, dbTransaction.ID)
	// todo check time
	require.Equal(t, amount, dbTransaction.Amount)
	require.Equal(t, tr.Title, dbTransaction.Title)
	require.Equal(t, tr.Category, dbTransaction.Category)

	err = r.db.Delete(dbTransaction).Error
	require.Nil(t, err)
}

func TestTransactionsRepository_DeleteTransaction(t *testing.T) {
	// todo add test
}

func TestTransactionsRepository_GetTransactionsByTimeRange(t *testing.T) {
	// todo add test
}

func TestTransactionsRepository_GetTransactionsByTimeRangeCategories(t *testing.T) {
	// todo add test
}
