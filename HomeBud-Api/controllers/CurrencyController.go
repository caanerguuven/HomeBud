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

//CreateCurrency function
func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	_cur := &modelOf.Currency{}
	utils.ParseBody(r, _cur)
	result := _cur.CreateCurrency()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteCurrency function
func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyID := vars["Id"]
	id, err := strconv.ParseInt(currencyID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	currency := modelOf.DeleteCurrency(id)
	response, _ := json.Marshal(currency)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllCurrencies function
func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies := modelOf.GetAllCurrencies()
	response, _ := json.Marshal(currencies)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetCurrencyByID function
func GetCurrencyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyID := vars["Id"]
	id, err := strconv.ParseInt(currencyID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	currency, _ := modelOf.GetCurrencyByID(id)
	response, _ := json.Marshal(&currency)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateCurrency function
func UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	var currencyDto = &modelOf.Currency{}
	utils.ParseBody(r, currencyDto)
	vars := mux.Vars(r)
	currencyID := vars["Id"]
	id, err := strconv.ParseInt(currencyID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	currency, db := modelOf.GetCurrencyByID(id)

	if currencyDto.Name != "" {
		currency.Name = currencyDto.Name
	}

	if currencyDto.Active != nil {
		currency.Active = currencyDto.Active
	}

	db.Save(&currency)
	response, _ := json.Marshal(&currency)

	utils.WriteResponse(w, http.StatusOK, response)
}
