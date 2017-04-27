package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func TestGetTransactionByIDHandler_ServeHTTP_Success(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	user, _ := c.CreateUser()
	account, _ := c.CreateAccount(user.ID, "test", iso4217.USD)
	transaction, _ := c.CreateTransaction(account.ID, time.Now(), 5, "title 1", "category 1")
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewGetTransactionByIDHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}/transaction/{transaction_id}", handler).Methods(http.MethodGet)

	request := httptest.NewRequest(http.MethodGet, "/account/1/transaction/1", nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, http.StatusOK, response.Code)

	transaction, _ = c.GetTransactionByID(transaction.ID)
	require.Equal(t, float64(5), transaction.Amount)
	require.Equal(t, "title 1", transaction.Title)
	require.Equal(t, "category 1", transaction.Category)
}
