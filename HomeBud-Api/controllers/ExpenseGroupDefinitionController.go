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

//CreateExpenseGroupDefinition function
func CreateExpenseGroupDefinition(w http.ResponseWriter, r *http.Request) {
	_exgd := &modelOf.ExpenseGroupDefinition{}
	utils.ParseBody(r, _exgd)
	result := _exgd.CreateExpenseGroupDefinition()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteExpenseGroupDefinition function
func DeleteExpenseGroupDefinition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	expenseGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseGroupDefinition := modelOf.DeleteExpenseGroupDefinition(id)
	response, _ := json.Marshal(expenseGroupDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllExpenseGroupDefinitions function
func GetAllExpenseGroupDefinitions(w http.ResponseWriter, r *http.Request) {
	expenseGroupDefinitions := modelOf.GetAllExpenseGroupDefinitions()
	response, _ := json.Marshal(expenseGroupDefinitions)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetExpenseGroupDefinitionByID function
func GetExpenseGroupDefinitionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	expenseGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseGroupDefinition, _ := modelOf.GetExpenseGroupDefinitionByID(id)
	response, _ := json.Marshal(&expenseGroupDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateExpenseGroupDefinition function
func UpdateExpenseGroupDefinition(w http.ResponseWriter, r *http.Request) {
	var expenseGroupDefinitionDto = &modelOf.ExpenseGroupDefinition{}
	utils.ParseBody(r, expenseGroupDefinitionDto)
	vars := mux.Vars(r)
	expenseGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseGroupDefinition, db := modelOf.GetExpenseGroupDefinitionByID(id)
	if expenseGroupDefinitionDto.Name != "" {
		expenseGroupDefinition.Name = expenseGroupDefinitionDto.Name
	}

	if expenseGroupDefinitionDto.Active != nil {
		expenseGroupDefinition.Active = expenseGroupDefinitionDto.Active
	}

	db.Save(&expenseGroupDefinition)
	response, _ := json.Marshal(&expenseGroupDefinition)

	utils.WriteResponse(w, http.StatusOK, response)
}
