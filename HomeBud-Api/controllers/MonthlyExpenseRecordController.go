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

//CreateMonthlyExpenseRecord function
func CreateMonthlyExpenseRecord(w http.ResponseWriter, r *http.Request) {
	_moer := &modelOf.MonthlyExpenseRecord{}
	utils.ParseBody(r, _moer)
	result := _moer.CreateMonthlyExpenseRecord()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteMonthlyExpenseRecord function
func DeleteMonthlyExpenseRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyExpenseRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyExpenseRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyExpenseRecord := modelOf.DeleteMonthlyExpenseRecord(id)
	response, _ := json.Marshal(monthlyExpenseRecord)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllMonthlyExpenseRecords function
func GetAllMonthlyExpenseRecords(w http.ResponseWriter, r *http.Request) {
	monthlyExpenseRecords := modelOf.GetAllMonthlyExpenseRecords()
	response, _ := json.Marshal(monthlyExpenseRecords)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetMonthlyExpenseRecordByID function
func GetMonthlyExpenseRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyExpenseRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyExpenseRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyExpenseRecord, _ := modelOf.GetMonthlyExpenseRecordByID(id)
	response, _ := json.Marshal(&monthlyExpenseRecord)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateMonthlyExpenseRecord function
func UpdateMonthlyExpenseRecord(w http.ResponseWriter, r *http.Request) {
	var monthlyExpenseRecordDto = &modelOf.MonthlyExpenseRecord{}
	utils.ParseBody(r, monthlyExpenseRecordDto)
	vars := mux.Vars(r)
	monthlyExpenseRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyExpenseRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyExpenseRecord, db := modelOf.GetMonthlyExpenseRecordByID(id)
	if monthlyExpenseRecordDto.Year != nil {
		monthlyExpenseRecord.Year = monthlyExpenseRecordDto.Year
	}

	if monthlyExpenseRecordDto.Month != nil {
		monthlyExpenseRecord.Month = monthlyExpenseRecordDto.Month
	}

	if monthlyExpenseRecordDto.ExpenseDefinitionID != nil {
		monthlyExpenseRecord.ExpenseDefinitionID = monthlyExpenseRecordDto.ExpenseDefinitionID
	}

	if monthlyExpenseRecordDto.EstimatedAmount != nil {
		monthlyExpenseRecord.EstimatedAmount = monthlyExpenseRecordDto.EstimatedAmount
	}

	if monthlyExpenseRecordDto.ActualAmount != nil {
		monthlyExpenseRecord.ActualAmount = monthlyExpenseRecordDto.ActualAmount
	}

	if monthlyExpenseRecordDto.CurrencyID != nil {
		monthlyExpenseRecord.CurrencyID = monthlyExpenseRecordDto.CurrencyID
	}

	if monthlyExpenseRecordDto.Description != "" {
		monthlyExpenseRecord.Description = monthlyExpenseRecordDto.Description
	}

	db.Save(&monthlyExpenseRecord)
	response, _ := json.Marshal(&monthlyExpenseRecord)

	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdatePaidStatus function
func UpdatePaidStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyExpenseRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyExpenseRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	isPaid := vars["isPaid"]
	paidStatus, err := strconv.ParseBool(isPaid)
	if err != nil {
		fmt.Println("boolean parse error")
		return
	}

	monthlyExpenseRecord, db := modelOf.GetMonthlyExpenseRecordByID(id)
	monthlyExpenseRecord.IsPaid = paidStatus

	db.Save(&monthlyExpenseRecord)
	response, _ := json.Marshal(&monthlyExpenseRecord)

	utils.WriteResponse(w, http.StatusOK, response)
}
