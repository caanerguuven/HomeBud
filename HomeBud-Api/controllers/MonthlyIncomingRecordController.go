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

//CreateMonthlyIncomingRecord function
func CreateMonthlyIncomingRecord(w http.ResponseWriter, r *http.Request) {
	_moir := &modelOf.MonthlyIncomingRecord{}
	utils.ParseBody(r, _moir)
	result := _moir.CreateMonthlyIncomingRecord()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteMonthlyIncomingRecord function
func DeleteMonthlyIncomingRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyIncomingRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyIncomingRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyIncomingRecord := modelOf.DeleteMonthlyIncomingRecord(id)
	response, _ := json.Marshal(monthlyIncomingRecord)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllMonthlyIncomingRecords function
func GetAllMonthlyIncomingRecords(w http.ResponseWriter, r *http.Request) {
	monthlyIncomingRecords := modelOf.GetAllMonthlyIncomingRecords()
	response, _ := json.Marshal(monthlyIncomingRecords)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetMonthlyIncomingRecordByID function
func GetMonthlyIncomingRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyIncomingRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyIncomingRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyIncomingRecord, _ := modelOf.GetMonthlyIncomingRecordByID(id)
	response, _ := json.Marshal(&monthlyIncomingRecord)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateMonthlyIncomingRecord function
func UpdateMonthlyIncomingRecord(w http.ResponseWriter, r *http.Request) {
	var monthlyIncomingRecordDto = &modelOf.MonthlyIncomingRecord{}
	utils.ParseBody(r, monthlyIncomingRecordDto)
	vars := mux.Vars(r)
	monthlyIncomingRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyIncomingRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	monthlyIncomingRecord, db := modelOf.GetMonthlyIncomingRecordByID(id)
	if monthlyIncomingRecordDto.Year != nil {
		monthlyIncomingRecord.Year = monthlyIncomingRecordDto.Year
	}

	if monthlyIncomingRecordDto.Month != nil {
		monthlyIncomingRecord.Month = monthlyIncomingRecordDto.Month
	}

	if monthlyIncomingRecordDto.IncomingDefinitionID != nil {
		monthlyIncomingRecord.IncomingDefinitionID = monthlyIncomingRecordDto.IncomingDefinitionID
	}

	if monthlyIncomingRecordDto.EstimatedAmount != nil {
		monthlyIncomingRecord.EstimatedAmount = monthlyIncomingRecordDto.EstimatedAmount
	}

	if monthlyIncomingRecordDto.ActualAmount != nil {
		monthlyIncomingRecord.ActualAmount = monthlyIncomingRecordDto.ActualAmount
	}

	if monthlyIncomingRecordDto.CurrencyID != nil {
		monthlyIncomingRecord.CurrencyID = monthlyIncomingRecordDto.CurrencyID
	}

	if monthlyIncomingRecordDto.Description != "" {
		monthlyIncomingRecord.Description = monthlyIncomingRecordDto.Description
	}

	db.Save(&monthlyIncomingRecord)
	response, _ := json.Marshal(&monthlyIncomingRecord)

	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateTakenStatus function
func UpdateTakenStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monthlyIncomingRecordID := vars["Id"]
	id, err := strconv.ParseInt(monthlyIncomingRecordID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	isTaken := vars["isTaken"]
	takenStatus, err := strconv.ParseBool(isTaken)
	if err != nil {
		fmt.Println("boolean parse error")
		return
	}

	monthlyIncomingRecord, db := modelOf.GetMonthlyIncomingRecordByID(id)
	monthlyIncomingRecord.IsTaken = takenStatus

	db.Save(&monthlyIncomingRecord)
	response, _ := json.Marshal(&monthlyIncomingRecord)

	utils.WriteResponse(w, http.StatusOK, response)
}
