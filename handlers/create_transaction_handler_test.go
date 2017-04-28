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
	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func TestCreateTransactionHandler_ServeHTTP_Success(t *testing.T) {
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
	log := logrus.New()
	log.Out = ioutil.Discard

	handler := NewCreateTransactionHandler(c, log)
	router := mux.NewRouter()
	router.Handle(URLRouteCreateTransaction, handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "10")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, BuildPathCreateTransaction(user.UUID, account.UUID), strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	require.Equal(t, http.StatusOK, response.Code)
}
