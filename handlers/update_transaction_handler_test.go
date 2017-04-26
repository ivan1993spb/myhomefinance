package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func TestUpdateTransactionHandler_ServeHTTP_Success(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	user, _ := c.CreateUser()
	account, _ := c.CreateAccount(user.ID, "test", iso4217.USD)
	transaction, _ := c.CreateTransaction(account.ID, time.Now(), 5, "title 1", "category 1")
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewUpdateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}/transaction/{transaction_id}", handler).Methods(http.MethodPut)

	data := &url.Values{}
	data.Add(fieldAmount, "22")
	data.Add(fieldTime, time.Now().Format(apiDateFormat))
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPut, "/account/1/transaction/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, http.StatusOK, response.Code)

	transaction, _ = c.GetTransactionByID(transaction.ID)
	require.Equal(t, float64(22), transaction.Amount)
	require.Equal(t, "title", transaction.Title)
	require.Equal(t, "category", transaction.Category)
}
