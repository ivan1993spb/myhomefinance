package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
)

func TestCreateTransactionHandler_ServeHTTP_EmptyAccountID(t *testing.T) {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	handler := NewCreateTransactionHandler(c, logrus.New())
	router := mux.NewRouter()
	router.Handle("/account/{account_id}", handler).Methods(http.MethodPost)

	data := &url.Values{}
	data.Add(fieldAmount, "22")
	data.Add(fieldTime, "2006-Jan-02")
	data.Add(fieldTitle, "title")
	data.Add(fieldCategory, "category")
	request := httptest.NewRequest(http.MethodPost, "/account/4", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	fmt.Printf("%+v", transactionRepository)
}
