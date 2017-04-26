package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func TestCreateTransactionHandler_ServeHTTP_EmptyAccountID(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "22")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusNotFound)
}

func TestCreateTransactionHandler_ServeHTTP_InvalidAccountID(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "22")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/invalid", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestCreateTransactionHandler_ServeHTTP_EmptyAmount(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestCreateTransactionHandler_ServeHTTP_InvalidAmount(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "invalid")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestCreateTransactionHandler_ServeHTTP_EmptyTime(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "10")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestCreateTransactionHandler_ServeHTTP_InvalidTime(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "10")
	data.Add(fieldTime, "invalid")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusBadRequest)
}

func TestCreateTransactionHandler_ServeHTTP_Success(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()
	log.Out = ioutil.Discard
	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "10")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/1", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, response.Code, http.StatusOK)
}
