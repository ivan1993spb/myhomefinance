package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
	"github.com/ivan1993spb/myhomefinance/models"
)

func TestDeleteTransactionHandler_ServeHTTP_EmptyTransactionID(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewDeleteTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}/transaction/{transaction_id}", handler).Methods(http.MethodDelete)

	request := httptest.NewRequest(http.MethodDelete, "/account/1/transaction/", nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusNotFound)
}

func TestDeleteTransactionHandler_ServeHTTP_InvalidTransactionID(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewDeleteTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}/transaction/{transaction_id}", handler).Methods(http.MethodDelete)

	request := httptest.NewRequest(http.MethodDelete, "/account/1/transaction/invalid", nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestDeleteTransactionHandler_ServeHTTP_Success(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewDeleteTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}/transaction/{transaction_id}", handler).Methods(http.MethodDelete)

	request := httptest.NewRequest(http.MethodDelete, "/account/1/transaction/1", nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusOK)
}
