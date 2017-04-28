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

func TestDeleteTransactionHandler_ServeHTTP_Success(t *testing.T) {
	userRepository, err := memoryrepository.NewUserRepository()
	require.Nil(t, err)
	accountRepository, err := memoryrepository.NewAccountRepository()
	require.Nil(t, err)
	transactionRepository, err := memoryrepository.NewTransactionRepository()
	require.Nil(t, err)
	c := core.New(userRepository, accountRepository, transactionRepository)
	user, err := c.CreateUser()
	require.Nil(t, err)
	account, err := c.CreateAccount(user.UUID, "account", iso4217.USD)
	require.Nil(t, err)
	transaction, err := c.CreateTransaction(user.UUID, account.UUID, time.Unix(1, 0), 5, "title 1", "category 1")
	require.Nil(t, err)
	log := logrus.New()
	log.Out = ioutil.Discard

	handler := NewDeleteTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle(URLRouteDeleteTransaction, handler).Methods(http.MethodDelete)

	request := httptest.NewRequest(http.MethodDelete, BuildPathDeleteTransaction(user.UUID, account.UUID, transaction.UUID), nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, http.StatusOK, response.Code)
}
