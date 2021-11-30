package controllers

import (
	"encoding/json"
	"fmt"
	modelOf "homebud/models"
	utils "homebud/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateAccount function
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	_cur := &modelOf.Account{}
	utils.ParseBody(r, _cur)
	result := _cur.CreateAccount()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteAccount function
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["Id"]
	id, err := strconv.ParseInt(accountID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	account := modelOf.DeleteAccount(id)
	response, _ := json.Marshal(account)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllAccounts function
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := modelOf.GetAllAccounts()
	response, _ := json.Marshal(accounts)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAccountByID function
func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["Id"]
	id, err := strconv.ParseInt(accountID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	account, _ := modelOf.GetAccountByID(id)
	response, _ := json.Marshal(&account)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateAccount function
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var AccountDto = &modelOf.Account{}
	utils.ParseBody(r, AccountDto)
	vars := mux.Vars(r)
	accountID := vars["Id"]
	id, err := strconv.ParseInt(accountID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	account, db := modelOf.GetAccountByID(id)
	if AccountDto.UserName != "" {
		account.UserName = AccountDto.UserName
	}

	if AccountDto.Password != "" {
		account.Password = AccountDto.Password
	}

	if AccountDto.FirstName != "" {
		account.FirstName = AccountDto.FirstName
	}

	if AccountDto.LastName != "" {
		account.LastName = AccountDto.LastName
	}

	if AccountDto.Active != nil {
		account.Active = AccountDto.Active
	}

	db.Save(&account)
	response, _ := json.Marshal(&account)

	utils.WriteResponse(w, http.StatusOK, response)
}
